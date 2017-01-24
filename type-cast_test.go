package Figo

import (
	"github.com/quexer/utee"
	"log"
	"testing"
)

func TestTpInt(t *testing.T) {
	v, err := TpInt(int64(1000))
	utee.Chk(err)
	log.Println(v == 1000)
	v, err = TpInt("1024")
	utee.Chk(err)
	log.Println(v == 1024)
	v, err = TpInt([]byte("768"))
	utee.Chk(err)
	log.Println(v == 768)
}

func TestTpTpInt64(t *testing.T) {
	v, err := TpInt64(int64(1000))
	utee.Chk(err)
	log.Println(v == 1000)
	v, err = TpInt64("1024")
	utee.Chk(err)
	log.Println(v == 1024)
	v, err = TpInt64([]byte("768"))
	utee.Chk(err)
	log.Println(v == 768)
}
