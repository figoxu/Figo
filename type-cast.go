package Figo

import (
	"fmt"
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
	return 0, fmt.Errorf("unexpected type for Int, got type %T", v)
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
