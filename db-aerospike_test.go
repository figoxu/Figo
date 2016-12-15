package Figo

import (
	"log"
	"testing"
)

func TestAs(t *testing.T) {
	ac := AsUtee.AsConnect("127.0.0.1:3000")
	setInfo := AsUtee.NewSetInfo("push", "test")
	type Val struct {
		Offset int64
	}
	err := AsUtee.Put(ac, setInfo, "hello", &Val{1027})
	log.Println("Put Err @err:", err)
	var val Val
	err = AsUtee.Get(ac, setInfo, "hello", &val)
	log.Println("Get Err @err:", err)
	log.Println(val)
}
