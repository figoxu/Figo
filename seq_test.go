package Figo

import (
	"log"
	"testing"
)

func TestSeqRedisNext(t *testing.T) {
	seqRedis := SeqRedis{
		rp:  RedisPool("localhost:6379", ""),
		key: "test_seq",
	}
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
