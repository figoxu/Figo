package Figo

import (
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
