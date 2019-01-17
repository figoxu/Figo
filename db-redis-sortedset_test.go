package Figo

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestRedisSortedSet_ZRangeByScore(t *testing.T) {
	rp := RedisPool("127.0.0.1:6379", "")
	sortedSet := NewRedisSortedSet("ss_test", rp)
	nanoSec := time.Now().UnixNano()
	add := func() {
		for i := 0; i < 1000; i++ {
			sortedSet.ZAdd(nanoSec+int64(i), fmt.Sprint("testKey", i))
		}
	}
	list := func() []string {
		skeys := sortedSet.ZRangeByScore(0, nanoSec+10000)
		for _, skey := range skeys {
			log.Println(skey)
		}
		return skeys
	}
	rm := func(skeys ...string) {
		for _, skey := range skeys {
			sortedSet.ZRem(skey)
		}
	}
	log.Println("Test>Add:")
	add()
	log.Println("List")
	vs := list()
	log.Println("Test>Rm:")
	rm(vs...)
	log.Println("After>Rm:")
	list()
}
