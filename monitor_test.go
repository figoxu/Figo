package Figo

import (
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
	m.Post("/test/post", MonitorMid, simpleHandle)
	m.Get("/test/get", MonitorMid, simpleHandle)
	http.Handle("/", m)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	time.Sleep(time.Second * time.Duration(2))
	cb := func(msgs ...string) {
		log.Println(msgs)
		log.Println(msgs)
	}
	MonitorCall("http://localhost:8080/test/post", "post", cb)
	MonitorCall("http://localhost:8080/test/get", "get", cb)
}
