package Figo

import (
	"testing"
	"log"
	"fmt"
	"time"
)

func TestBatchMemQueue_DeqN(t *testing.T) {
	q:=NewBatchMemQueue(10e5,16,10, func(vs []interface{}) {
		if len(vs)>0 {
			log.Println(len(vs))
		}
	})
	for i:=0;i<103;i++ {
		q.Enq(fmt.Sprint(i))
	}
	time.Sleep(time.Second*time.Duration(5))
}

func TestBatchMemQueue_DeqOne(t *testing.T) {
	q:=NewBatchMemQueue(10e5,16,1, func(vs []interface{}) {
		for _,v:=range vs {
			log.Println(v)
		}
	})
	for i:=0;i<103;i++ {
		q.Enq(fmt.Sprint(i))
	}
	time.Sleep(time.Second*time.Duration(5))
}

