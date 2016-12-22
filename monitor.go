package Figo

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/go-martini/martini"
	"github.com/pborman/uuid"
	"github.com/quexer/utee"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	MONITOR_HB_KEY = "MONITOR-HEARTBEAT"
	MONITOR_HB_VAL = "ping"
	MONITOR_HB_RSP = "pong"
	MONITOR_CB_KEY = "MONITOR-CALLBACK"
	MONITOR_CB_VAL = "q1w2e3r4t5"
)

var (
	MONITOR_BODY = url.Values{
		"sys_monitor_body": []string{"handle the case of skip body by server"},
	}
)

func MonitorMidCheck(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(MONITOR_HB_KEY) == MONITOR_HB_VAL {
		http.Error(w, MONITOR_HB_RSP, 200)
		return
	}
}

func MonitorCall(restApi, method string, warn func(...string)) {
	var b []byte
	var err error
	header := make(http.Header)
	header.Add(MONITOR_HB_KEY, MONITOR_HB_VAL)
	if "GET" == strings.ToUpper(method) {
		b, err = HttpGet(restApi, MONITOR_BODY, header)
	} else {
		b, err = HttpPost(restApi, MONITOR_BODY, header)
	}
	if err != nil {
		warn("Service Has Http Error @restApi:", restApi, " @rsp:", string(b))
		return
	}
	if !strings.Contains(string(b), MONITOR_HB_RSP) {
		warn("Service Has Check Error @restApi:", restApi, " @rsp:", string(b))
		return
	}
	log.Println("Exam At @restApi:", restApi, " @method:", method, " @rsp:", string(b))
}

func HttpPost(api string, q url.Values, header http.Header) ([]byte, error) {
	return HttpRequest(api, "POST", header, q)
}

func HttpGet(api string, q url.Values, header http.Header) ([]byte, error) {
	return HttpRequest(api, "GET", header, q)
}

func HttpRequest(api, method string, header http.Header, q url.Values) ([]byte, error) {
	method = strings.ToUpper(method)
	var req *http.Request
	var err error
	if q != nil {
		sreader := strings.NewReader(q.Encode())
		req, err = http.NewRequest(method, api, sreader)
	} else {
		req, err = http.NewRequest(method, api, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("[http] err %s, %s\n", api, err)
	}
	client := http.DefaultClient
	for k, vs := range header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[http] err %s, %s\n", api, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("[http get] status err %s, %d\n", api, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func MonitorMidCB(w http.ResponseWriter, r *http.Request) {
	getV := func(key string, header http.Header) string {
		uKey, lKey := strings.ToUpper(key), strings.ToLower(key)
		if v := header.Get(key); v != "" {
			log.Println("@key:", key, " @v:", v)
			return v
		}
		if v := header.Get(uKey); v != "" {
			log.Println("@key:", uKey, " @v:", v)
			return v
		}
		if v := header.Get(lKey); v != "" {
			log.Println("@key:", lKey, " @v:", v)
			return v
		}
		return ""
	}
	if getV(MONITOR_HB_KEY, r.Header) == MONITOR_HB_VAL {
		if api := getV(MONITOR_CB_KEY, r.Header); api != "" {
			HttpGet(api, MONITOR_BODY, make(http.Header))
		}
		return
	}
}

type MonitorCallBack struct {
	tc    *utee.TimerCache
	cbURL string
	warn  func(...string)
}

func NewMonitorCallBack(cbURL string, ttl int, warn func(...string)) *MonitorCallBack {
	checkFail := func(key, value interface{}) {
		defer Catch()
		val, err := TpString(value)
		utee.Chk(err)
		warn("Service Has Http Error @restApi:", val)
	}
	return &MonitorCallBack{
		tc:    utee.NewTimerCache(ttl, checkFail),
		cbURL: cbURL,
		warn:  warn,
	}
}

func (p *MonitorCallBack) Handler() func(martini.Params) (int, string) {
	handle := func(params martini.Params) (int, string) {
		p.tc.Remove(params["id"])
		return 200, "ok"
	}
	return handle
}

func (p *MonitorCallBack) BeegoHandler() func(c *context.Context) {
	return func(c *context.Context) {
		id := c.Input.Param(":id")
		p.tc.Remove(id)
	}
}

func (p *MonitorCallBack) CallOnTime(cronExp, restApi, method string, warn func(...string)) {
	c := cron.New()
	c.AddFunc(cronExp, func() {
		log.Println("invoke @api:", restApi, "@method:", method)
		header := make(http.Header)
		header.Add(MONITOR_HB_KEY, MONITOR_HB_VAL)
		id := uuid.NewUUID().String()
		api := strings.Replace(p.cbURL, ":id", id, -1)
		header.Add(MONITOR_CB_KEY, api)
		p.tc.Put(id, restApi)
		var b []byte
		var err error
		if "GET" == strings.ToUpper(method) {
			b, err = HttpGet(restApi, MONITOR_BODY, header)
		} else {
			b, err = HttpPost(restApi, MONITOR_BODY, header)
		}
		log.Println("[monitor] @rsp:", string(b), "  @err:", err)

	})
	c.Start()
}
