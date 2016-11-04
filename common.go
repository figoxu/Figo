package Figo

import (
	"log"
	"runtime/debug"
)

func Catch() {
	if err := recover(); err != nil {
		log.Println(string(debug.Stack()))
		log.Println(err, " (recover)")
	}
}

func Clone(src interface{}) interface{} {
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
