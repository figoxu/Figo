package Figo

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/quexer/utee"
	"log"
	"os"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
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
		err = business()
	}
	if retry > 0 && tips != "" {
		success := (err == nil)
		log.Println(tips, " Execute With ", retry, " times .  @SuccessFlag:", success)
	}
}

func ParseUrl(s string) (string, int, error) {
	a := strings.Split(s, ":")
	if len(a) != 2 {
		return "", 0, fmt.Errorf("bad url %s", s)
	}
	port, err := strconv.Atoi(a[1])
	return a[0], port, err
}

func ReadInput(tips string) string {
	log.Println(tips)
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	return string(data)
}
