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

func RetryExe(business func() error, times int, tips string) {
	err := business()
	retry := 0
	for err != nil && retry < times {
		retry++
		if err = business(); err == nil {
			break
		}
	}
	if retry > 0 && tips != "" {
		success := (err == nil)
		log.Println(tips, " Execute With ", retry, " times .  @SuccessFlag:", success)
	}
}
