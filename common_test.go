package Figo

import (
	"github.com/gogap/errors"
	"github.com/quexer/utee"
	"log"
	"testing"
)

func TestCatch(t *testing.T) {
	log.Println("Before")
	catchSampleMethod()
	log.Println("After")
}

func catchSampleMethod() {
	defer Catch()
	log.Println("Hello")
	utee.Chk(errors.New("Some Error"))
}
