package Figo

import (
	"testing"
	"log"
)

func TestNewRedisHash(t *testing.T) {
	rp := RedisPool("localhost:6379", "")
	hash:=NewRedisHash(rp,"FIGO_TEST_HASH_KEY")
	hash.Set(Skv{
		K:"Hello",
		V:"World",
	})
	log.Println(hash.Get("Hello"))

	hash.Set([]Skv{Skv{
		"a","1",
	},Skv{
		"b","2",
	},Skv{
		"c","3",
	},Skv{
		"d","4",
	},Skv{
		"e","5",
	},Skv{
		"f","6",
	}}...)
	log.Println(hash.Get("a","b","c","d","e","f"))
	log.Println(hash.Get("KEY_NEVER_EXIST","Hello","a","b","c","d","e","f"))

}

