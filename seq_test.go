package Figo

import (
	"log"
	"testing"
)

func TestSeqRedisNext(t *testing.T) {
	seqRedis := NewSeqRedis(RedisPool("localhost:6379", ""), "test_seq", -8)
	for i := 0; i < 10; i++ {
		log.Println(seqRedis.Next())
	}
}

func TestSeqMemNext(t *testing.T) {
	seqMem := SeqMem{}
	for i := 0; i < 10; i++ {
		log.Println(seqMem.Next())
	}
}
