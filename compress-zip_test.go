package Figo

import (
	"github.com/quexer/utee"
	"testing"
)

func TestZip(t *testing.T) {
	dir := "/home/figo/develop/env/GOPATH/src/github.com/figoxu/"
	err := Zip(dir, "/home/figo/delete/zipResult/figo.zip")
	utee.Chk(err)
}

func TestUnZip(t *testing.T) {
	UnZip("/home/figo/delete/zipResult/figo.zip", "/home/figo/delete/unzipResult/")
}
