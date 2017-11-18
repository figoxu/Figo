package Figo

import (
	"github.com/quexer/utee"
	"log"
	"testing"
)

func TestNewAesHelp(t *testing.T) {
	pwd := []byte{208, 130, 159, 99, 231, 235, 183, 11, 143, 170, 60, 9, 183, 15, 29, 171}
	aesHelp := NewAesHelp(pwd)
	bs, err := aesHelp.Encrypt([]byte("HelloWorld世界你好안녕하세요."))
	utee.Chk(err)
	bs, err = aesHelp.Decrypt(bs)
	log.Println(string(bs))
}
