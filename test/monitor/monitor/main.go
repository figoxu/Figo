package main

import (
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Handlers(martini.Recovery())
	warn := func(msgs ...string) {
		str := fmt.Sprint(msgs)
		client := Figo.GetSMTPClient("xujh945@qq.com", "xxxxx", "smtp.qq.com:25")
		client.Send("xujh945@qq.com", str, str)
		client.Send("jianhui.xu@tendcloud.com", str, str)
	}
	mcb := Figo.NewMonitorCallBack("http://localhost:9090/cb/:id", 10, warn)
	m.Get("/cb/:id", mcb.Handler())
	http.Handle("/", m)
	mcb.CallOnTime("0/5 * * * * ?", "http://localhost:8080/test/post", "POST", warn)
	mcb.CallOnTime("0/5 * * * * ?", "http://localhost:8080/test/get", "GET", warn)
	mcb.CallOnTime("0/5 * * * * ?", "http://localhost:8080/test/get/withOutMonitor", "GET", warn)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
