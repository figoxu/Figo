package Figo

import (
	"log"
	"testing"
	"time"
)

func TestNewTimerCache(t *testing.T) {
	tc := NewTimerCache(10, func(key, value interface{}) {
		log.Println("expire @key:", key, "  @value:", value)
	})
	tc.Put("foo", "bar")
	tc.Put("hello", "world")
	log.Println(tc.Get("foo"))
	log.Println(tc.Get("hello"))
	time.Sleep(time.Duration(11) * time.Second)
}
