package Figo

import (
	"github.com/quexer/utee"
	"log"
	"testing"
)

func TestRedisSetAndGet(t *testing.T) {
	rp := RedisPool("localhost:6379", "")
	RedisSet(rp, "FOO", "BAR")
	v, err := RedisGet(rp, "FOO")
	utee.Chk(err)
	log.Println(TpString(v))
}
