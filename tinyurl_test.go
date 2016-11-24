package Figo

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"testing"
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
