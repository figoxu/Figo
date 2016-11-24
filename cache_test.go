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

func TestNewRedisCache(t *testing.T) {
	rc := NewRedisCache(RedisPool("localhost:6379", ""))
	rc.Put("foo", "bar")
	rc.Put("hello", "world")
	log.Println(TpString(rc.Get("foo")))
	log.Println(TpString(rc.Get("hello")))
}

func TestNewRedisTimerCache(t *testing.T) {
	rc := NewRedisTimerCache(RedisPool("localhost:6379", ""), 10)
	rc.Put("foo", "bar")
	rc.Put("hello", "world")
	log.Println(TpString(rc.Get("foo")))
	log.Println(TpString(rc.Get("hello")))
	time.Sleep(time.Duration(11) * time.Second)
	log.Println(TpString(rc.Get("foo")))
	log.Println(TpString(rc.Get("hello")))
}
