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
