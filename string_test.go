package Figo

import (
	"log"
	"testing"
)

func TestSnakeString(t *testing.T) {
	log.Println(SnakeString("FigoHowAreYou"))
}

func TestCamelString(t *testing.T) {
	log.Println(CamelString("FigoHowAreYou"))
func TestMd5Shard(t *testing.T) {
	shardSize := 2
	log.Println(Md5Shard("helloFooBarWorld", shardSize))
}
