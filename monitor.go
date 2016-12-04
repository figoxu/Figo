package Figo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	MONITOR_HB_KEY = "MONITOR_HEARTBEAT"
	MONITOR_HB_VAL = "ping"
	MONITOR_HB_RSP = "pong"
	MONITOR_CB_KEY = "MONITOR_CALLBACK"
	MONITOR_CB_VAL = "q1w2e3r4t5"
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
		b, err = HttpGet(restApi, header)
	} else {
		b, err = HttpPost(restApi, url.Values{}, header)
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

func HttpGet(api string, header http.Header) ([]byte, error) {
	return HttpRequest(api, "GET", header, nil)
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
