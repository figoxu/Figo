package Figo

import (
	"errors"
	"github.com/quexer/utee"
	"log"
	"reflect"
	"runtime/debug"
)

func Catch() {
	if err := recover(); err != nil {
		log.Println(string(debug.Stack()))
		log.Println(err, " (recover)")
	}
}

func Clone(src interface{}) interface{} {
	if reflect.TypeOf(src).Kind().String() == "ptr" {
		utee.Chk(errors.New("Can Not Clone An Point"))
	}
	dst := (src)
	return dst
}

func Exist(expect interface{}, objs ...interface{}) bool {
	for _, v := range objs {
		if expect == v {
			return true
		}
	}
	return false
}
