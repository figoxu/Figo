package Figo

import "reflect"

func TypeOf(sample interface{}) reflect.Type {
	return reflect.ValueOf(sample).Type()
}
