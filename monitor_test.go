package Figo

import (
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestMonitorMid(t *testing.T) {
	m := martini.Classic()
	m.Handlers(martini.Recovery())
	simpleHandle := func() (int, string) {
		log.Println("Enter Handler")
		return 200, "hello"
	}
	m.Post("/test/post", MonitorMidCheck, simpleHandle)
	m.Get("/test/get", MonitorMidCheck, simpleHandle)
	m.Get("/test/get/withOutMonitor", simpleHandle)
	http.Handle("/", m)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	time.Sleep(time.Second * time.Duration(2))
	cb := func(msgs ...string) {
		str := fmt.Sprint(msgs)
		client := GetSMTPClient("xujh945@qq.com", "xxxxxxxx", "smtp.qq.com:25")
		client.Send("xujh945@qq.com", str, str)
		client.Send("jianhui.xu@tendcloud.com", str, str)
	}
	MonitorCall("http://localhost:8080/test/post", "post", cb)
	MonitorCall("http://localhost:8080/test/get", "get", cb)
	MonitorCall("http://localhost:8080/test/get/withOutMonitor", "get", cb)
}
