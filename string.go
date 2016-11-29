package Figo

import (
	"fmt"
	"github.com/quexer/utee"
	"gopkg.in/iconv.v1"
	"log"
	"unicode/utf8"
)

func DealStringMessy(result string) string {
	if !utf8.ValidString(result) {
		cd, err := iconv.Open("utf-8", "gbk")
		utee.Chk(err)
		defer cd.Close()
		result = cd.ConvString(result)
		log.Println("Parse From GBK To  UTF-8")
	}
	return result
}

func splitUTF8BOM(str string) string {
	b := []byte(str)
	if len(b) < 3 {
		return str
	}
	prefix := fmt.Sprintf("%X", b[0:3])
	if prefix == "EFBBBF" {
		return string(b[3:len(b)])
	}
	return str
}
