package Figo

import (
	"testing"
	"github.com/quexer/utee"
	"log"
)

func TestNewAesHelp(t *testing.T) {
	pwd:=[]byte{208,130,159,99,231,235,183,11,143,170,60,9,183,15,29,171}
	aesHelp:= NewAesHelp(pwd)
	bs,err:=aesHelp.encrypt([]byte("HelloWorld世界你好안녕하세요."))
	utee.Chk(err)
	bs,err =aesHelp.decrypt(bs)
	log.Println(string(bs))
}
