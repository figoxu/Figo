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

func TestRedisBitmap(t *testing.T) {
	rp := RedisPool("localhost:6379", "")
	err := RedisSetbit(rp, "FOO", 1)
	log.Println("set bit ", err)

	err = RedisSetbit(rp, "FOO1", 2)
	log.Println("set bit ", err)

	exist, err := RedisGetbit(rp, "FOO", 1)
	log.Println("get bit exist", exist, "err", err)

	c, err := RedisBitcount(rp, "FOO")
	log.Println("count bit ", c, "err", err)

	err = RedisBitop(rp, "or", "dest", "FOO", "FOO1")
	log.Println("bitop or ", "err", err)

	err = RedisBitop(rp, "and", "dest1", "FOO", "FOO1", "FOO2")
	log.Println("bitop and ", "err", err)
}


func Test_r_Rpush(t *testing.T){
	rp := RedisPool("localhost:6379", "")
	k := "figo"
	RedisRpush(rp, k,"hello")
	RedisRpush(rp, k,"world")
	RedisRpush(rp, k,"how")
	RedisRpush(rp, k,"r")
	RedisRpush(rp, k,"u")
	log.Println(RedisLlen(rp, k))
	v,err:=RedisLpop(rp, k)
	utee.Chk(err)
	log.Println(v)
	l,err:=RedisLlen(rp, k)
	log.Println("@len:",l)
	for i:=0;i<l;i++ {
		log.Println(RedisLindex(rp,k,i))
	}
}
