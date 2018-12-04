package Figo

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"testing"
	"fmt"
)

func TestConvert(t *testing.T) {
	seqRedis := &SeqRedis{
		rp:  RedisPool("localhost:6379", ""),
		key: "test_seq",
	}
	cacheObj := NewTimerCache(3600, nil)
	tinyURL := NewTinyUrl(cacheObj, seqRedis)
	key := tinyURL.Convert("http://xxiongdi.iteye.com")
	log.Println("@key:", key)
	m := martini.Classic()
	m.Get("/:key", tinyURL.Handler())
	http.Handle("/", m)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TestUrlAppendParam(t *testing.T) {
	v := UrlAppendParam("http://figoxu.me/welcome.jsp?auth=true&zoom=99&age=18&male=true", "token", "123456trewq4rfvCDE#")
	fmt.Print(v)
}

func TestUrlExistParam(t *testing.T) {
	exist := UrlExistParam("http://www.baidu.com/index.html?foo=bar&hello=world", "foo")
	fmt.Println(exist)
	exist = UrlExistParam("http://www.baidu.com/index.html?foo=bar&hello=world", "somethingelse")
	fmt.Println(exist)

}

func TestUrlRemoveParam(t *testing.T) {
	v := UrlRemoveParam("http://www.baidu.com/index.html?foo=bar&hello=world", "foo")
	fmt.Println(v)
	v = UrlRemoveParam("http://admin.dev.app.startrip.vip/sdzadmin/index.html?basic_pure_token=MzEyOjgxMTE2OWRhLTRhNjEtNGIzMS1iMWQ0LWU0MDZjYzFlYjY2NQ%3D%3D#/campsdz/index", "basic_pure_token")
	fmt.Println(v)
}
