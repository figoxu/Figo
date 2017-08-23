package Figo

import (
	"fmt"
	"github.com/quexer/utee"
	"regexp"
	"strconv"
	"strings"
)

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

func ParserWithPipe(content string, ps ...Parser) []string {
	exc := func(parser Parser, contents ...string) []string {
		var result []string
		for _, v := range contents {
			result = append(result, parser.Exe(v)...)
		}
		return result
	}
	result := []string{content}
	for _, parser := range ps {
		result = exc(parser, result...)
	}
	return result
}

func Md5ShardPiece(key string, piece int) int {
	shardVal, err := strconv.ParseUint(utee.PlainMd5(key)[16:32], 16, 0)
	utee.Chk(err)
	shardVal = shardVal % uint64(piece)
	return int(shardVal)
}

func Md5Shard(key string, piece int) string {
	return fmt.Sprint(key, "_", Md5ShardPiece(key, piece))
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func SnakeStrings(strs ...string) (result []string) {
	for _, str := range strs {
		s := SnakeString(str)
		result = append(result, s)
	}
	return result
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
