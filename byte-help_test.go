package Figo

import (
	"testing"
	"log"
)

func TestByteHelp_Append(t *testing.T) {
	bs:=Bh.Append([]byte{1,2,3,4,5},[]byte{6,7,8},[]byte{9,10,11,12,13})
	log.Println(bs)
	str := Bh.BStr(bs)
	log.Println(str)
}

func TestByteHelp_B2I16(t *testing.T) {
	v:=Bh.B2I16([]byte{1,22})
	log.Println(v)
	v=Bh.B2I16([]byte{1,0})
	log.Println(v)
}


func TestByteHelp_I162B(t *testing.T) {
	v:=Bh.I162B(int16(278))
	log.Println(Bh.BStr(v))
	v=Bh.I162B(int16(256))
	log.Println(Bh.BStr(v))
}