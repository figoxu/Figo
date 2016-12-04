package main

import (
	"github.com/figoxu/Figo"
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Handlers(martini.Recovery())
	simpleHandle := func() (int, string) {
		log.Println("Enter Handler")
		return 200, "hello"
	}
	m.Post("/test/post", Figo.MonitorMidCB, simpleHandle)
	m.Get("/test/get", Figo.MonitorMidCB, simpleHandle)
	m.Get("/test/get/withOutMonitor", simpleHandle)
	http.Handle("/", m)
	log.Println("start server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
