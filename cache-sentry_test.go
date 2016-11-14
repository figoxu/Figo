package Figo

import (
	"github.com/quexer/utee"
	"log"
	"testing"
)

func TestSentry(t *testing.T) {
	tc := utee.NewTimerCache(60, nil)
	put := func(key, val interface{}) {
		tc.Put(key, val)
	}
	get := func(key interface{}) interface{} {
		return tc.Get(key)
	}
	cacheObj := NewCacheObj(put, get)

	notify := func(key interface{}) error {
		log.Println("need to notify @key:", key, " to clear the cache")
		return nil
	}
	sentry := NewSentry(notify, cacheObj)
	sentry.Put("hello", "world")
	sentry.Put("hello", "Cool")
	v := sentry.Get("hello")
	log.Println(v)
}
