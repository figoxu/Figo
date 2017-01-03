package Figo

import (
	"log"
	"testing"
)

func TestMd5Shard(t *testing.T) {
	shardSize := 2
	log.Println(Md5Shard("helloFooBarWorld", shardSize))
}
