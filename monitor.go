package Figo

import (
	"fmt"
	"github.com/go-martini/martini"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	AUTH_KEY = "FigoMonitor"
	AUTH_VAL = "ping"
	AUTH_RSP = "pong"
)

func MonitorMid(w http.ResponseWriter, r *http.Request, c martini.Context) {
	if r.Header.Get(AUTH_KEY) == AUTH_VAL {
		http.Error(w, AUTH_RSP, 200)
		return
	}
}

func MonitorCall(restApi, method string, warn func(...string)) {
	var b []byte
	var err error
	header := make(http.Header)
	header.Add(AUTH_KEY, AUTH_VAL)
	if "GET" == strings.ToUpper(method) {
		b, err = HttpGet(restApi, header)
	} else {
		q := url.Values{}
		b, err = HttpPost(restApi, q, header)
	}
	if err != nil {
		warn("Service Has Http Error @restApi:", restApi, " @rsp:", string(b))
		return
	}

	if strings.Contains(string(b), AUTH_RSP) {
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
