package Figo

import (
	"testing"
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
	"log"
	"fmt"
	"time"
)

func TestRedisQueue_Enq(t *testing.T) {
	rp:=RedisPool("127.0.0.1:6379","")
	rq:=NewRedisQueue(rp,"test",3,func(v string,err error){
		defer Catch()
		if err!=redis.ErrNil {
			utee.Chk(err)
		}
		log.Println(v)
	})
	for i:=0;i<100;i++{
		rq.Enq(fmt.Sprint("Hello ",i))
	}
	time.Sleep(time.Second*30)
}