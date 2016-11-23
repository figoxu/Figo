package Figo

import (
	"github.com/go-martini/martini"
	"github.com/quexer/utee"
	"log"
	"net/http"
	"testing"
)

func TestConvert(t *testing.T) {
	seqRedis := &SeqRedis{
		rp:  RedisPool("localhost:6379", ""),
		key: "test_seq",
	}
	tc := utee.NewTimerCache(3600, nil)
	put := func(key, val interface{}) {
		tc.Put(key, val)
	}
	get := func(key interface{}) interface{} {
		return tc.Get(key)
	}
	cacheObj := NewCacheObj(put, get)
	tinyURL := NewTinyUrl(cacheObj, seqRedis)
	key := tinyURL.Convert("http://xxiongdi.iteye.com")
	m := martini.Classic()
	m.Get("/figoxu", tinyURL.Handler(key))
	http.Handle("/", m)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
