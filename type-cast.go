package Figo

import (
	"fmt"
	"github.com/quexer/utee"
	"strconv"
)

func TpInt(v interface{}) (int, error) {
	switch reply := v.(type) {
	case int64:
		x := int(reply)
		if int64(x) != reply {
			return 0, strconv.ErrRange
		}
		return x, nil
	case []byte:
		n, err := strconv.ParseInt(string(reply), 10, 0)
		return int(n), err
	default:
		return strconv.Atoi(fmt.Sprint(reply))
	}
}

func TpInt64(v interface{}) (int64, error) {
	switch reply := v.(type) {
	case int64:
		return reply, nil
	case []byte:
		n, err := strconv.ParseInt(string(reply), 10, 64)
		return n, err
	default:
		return strconv.ParseInt(fmt.Sprint(reply), 10, 64)
	}
}

func TpString(v interface{}) (string, error) {
	switch reply := v.(type) {
	case []byte:
		return string(reply), nil
	case string:
		return reply, nil
	case nil:
		return "", nil
	}
	return "", fmt.Errorf("unexpected type for String, got type %T", v)
}

func TpFloat64(v interface{}) (float64, error) {
	s, err := TpString(v)
	utee.Chk(err)
	return strconv.ParseFloat(s, 64)
}

var ICast = InterfaceCast{}

type InterfaceCast struct {
}

func (p *InterfaceCast) FromString(v string) interface{} {
	return interface{}(v)
}

func (p *InterfaceCast) FromInt(v int) interface{} {
	return interface{}(v)
}

func (p *InterfaceCast) FromStrings(data []string) []interface{} {
	s := make([]interface{}, len(data))
	for i, v := range data {
		s[i] = v
	}
	return s
}

func (p *InterfaceCast) FromInts(data []int) []interface{} {
	s := make([]interface{}, len(data))
	for i, v := range data {
		s[i] = v
	}
	return s
}
