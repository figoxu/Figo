package Figo

import "fmt"

func TpString(reply interface{}) (string, error) {
	switch reply := reply.(type) {
	case []byte:
		return string(reply), nil
	case string:
		return reply, nil
	case nil:
		return "", nil
	}
	return "", fmt.Errorf("redigo: unexpected type for String, got type %T", reply)
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
