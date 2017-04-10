package Figo

import (
	"log"
	"testing"
)

func TestNewRedisSortedSet(t *testing.T) {
	rp := RedisPool("localhost:6379", "")
	ss := NewRedisSortedSet("figoSortedSet", 100, rp)
	ss.ZAdd(1000, "figo")
	log.Println("@total:", ss.ZCount(0, 100000))
	log.Println("@score:", ss.ZScore("figo"))
}
