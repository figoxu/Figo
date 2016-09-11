package Figo

import (
	"github.com/figoxu/utee"
	"testing"
)

func TestNewConfKV(t *testing.T) {
	kv, err := NewConfKV("./conf.kv")
	utee.Chk(err)
	kv.Write("Test", "Hello", "World")
	kv.Write("Test", "How", "U")
	kv.Write("Test", "Cool", "Y")
	kv.Write("Test", "JJ", "G")
	kv.Write("Figo", "PWD", "123456")
	kv.Write("Figo", "AGE", "32")
	kv.Write("Figo", "GLENDAR", "Man")
	kv.Write("Figo", "Cool", "Sample")
	kv.Flush()
	kv.Write("How", "r", "u")
	kv.Flush()
	kv.WriteRecord("lang", map[string]string{
		"unix":       "0",
		"python":     "1",
		"go":         "2",
		"javascript": "3",
	})
	kv.Flush()
}
