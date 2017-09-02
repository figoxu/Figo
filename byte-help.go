package Figo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

var Bh  = ByteHelp{}

type ByteHelp struct {

}

func (p *ByteHelp) I162B(n int16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &n)
	return bytesBuffer.Bytes()
}

func (p *ByteHelp) B2I16(b []byte) int16 {
	bytesBuffer := bytes.NewBuffer(b)
	log.Println(b)
	var v int16
	binary.Read(bytesBuffer, binary.BigEndian, &v)
	return int16(v)
}

func (p *ByteHelp) BStr(bs []byte)string{
	v:=""
	for  _,b :=range bs  {
		s:=fmt.Sprint(uint8(b))
		if len(s)<2{
			s=fmt.Sprint("00",s)
		}else if len(s)<3{
			s=fmt.Sprint("0",s)
		}
		v = fmt.Sprint(v," ",s)
	}
	return v
}

func (p *ByteHelp) Append(bss ...[]byte)[]byte {
	out := []byte{}
	for _,bs := range bss {

		for _,b:=range bs{
			log.Println(b)
			out = append(out,b)
		}
	}
	return out
}
