package Figo

import (
	"github.com/quexer/utee"
	"log"
	"testing"
	"time"
)

func TestRedisSetAndGet(t *testing.T) {
	rp := RedisPool("localhost:6379", "")
	RedisSet(rp, "FOO", "BAR")
	v, err := RedisGet(rp, "FOO")
	utee.Chk(err)
	log.Println(TpString(v))
}

func TestRedisSetEx(t *testing.T) {
	rp := RedisPool("localhost:6379", "")
	RedisSetEx(rp, "FOO", "BAR", 10)
	v, err := RedisGet(rp, "FOO")
	utee.Chk(err)
	log.Println(TpString(v))
	time.Sleep(time.Duration(11) * time.Second)
	v, err = RedisGet(rp, "FOO")
	utee.Chk(err)
	log.Println(TpString(v))
}
