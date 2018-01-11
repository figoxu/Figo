package Figo

import (
	"testing"
	"fmt"
	"log"
)

func TestRedisQFlask(t *testing.T){
	rp := RedisPool("localhost:6379", "")
	qflask := NewRedisQFlask(rp,"testFigo1",10)
	for i:=0;i<100;i++ {
		qflask.Put(fmt.Sprint("GoodHealth",i))
	}
	for i,v:= range  qflask.Get() {
		log.Println("@index:",i," @value:",v)
	}

}

func TestRedisQMultiFlask(t *testing.T) {
	rp := RedisPool("localhost:6379", "")
	qflask := NewRedisQMultiFlask(rp,10)
	for i:=0;i<100;i++ {
		for j:=0;j<5;j++ {
			qflask.Put(fmt.Sprint("k",j),fmt.Sprint("GoodHealth",i))
		}
	}
	for i:=0;i<5;i++{
		for i,v:= range  qflask.Get(fmt.Sprint("k",i)) {
			log.Println("@index:",i," @value:",v)
		}
	}
}

