package Figo

import (
	"github.com/quexer/utee"
	//"log"
	"flag"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"testing"
)

func TestConvert(t *testing.T) {
	tc := utee.NewTimerCache(60, nil)
	put := func(key, val interface{}) {
		tc.Put(key, val)
	}
	get := func(key interface{}) interface{} {
		return tc.Get(key)
	}
	cacheObj := NewCacheObj(put, get)
	c := NewTinyUrl(cacheObj)
	k := c.Convert("http://www.baidu.com")
	addr := flag.String("p", ":8888", "address where the server listen on")
	m := martini.Classic()
	m.Get("/td", c.GetRedirectUrlHandler(k))

	http.Handle("/", m)

	log.Fatal(http.ListenAndServe(*addr, nil))

}
