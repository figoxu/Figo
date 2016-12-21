package Figo

import (
	"fmt"
	"github.com/quexer/utee"
	"gopkg.in/iconv.v1"
	"log"
	"regexp"
	"strings"
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

func SplitUTF8BOM(str string) string {
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

type Parser struct {
	PrepareReg []string
	ProcessReg []string
}

func (p *Parser) Exe(content string) []string {
	prep := func(reg string, contents ...string) []string {
		var result []string
		for _, content := range contents {
			rs := regexp.MustCompile(reg).FindAllString(content, -1)
			result = append(result, rs...)
		}
		return result
	}
	proc := func(reg string, contents ...string) []string {
		var result []string
		for _, content := range contents {
			rs := regexp.MustCompile(reg).ReplaceAllString(content, "")
			result = append(result, rs)
		}
		return result
	}
	result := []string{content}
	for _, reg := range p.PrepareReg {
		result = prep(reg, result...)
	}
	for _, reg := range p.ProcessReg {
		result = proc(reg, result...)
	}
	return TrimAndClear(result...)
}

func TrimAndClear(strs ...string) []string {
	result := []string{}
	for _, v := range strs {
		v = strings.TrimSpace(v)
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}
