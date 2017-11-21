package Figo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

var Bh = ByteHelp{}

type ByteHelp struct {
}

func (p *ByteHelp) Equal(bs1 ,bs2 []byte) bool {
	if len(bs1)!=len(bs2){
		return false
	}
	for index,v := range bs1 {
		if bs2[index]!=v {
			return false
		}
	}
	return true
}

func (p *ByteHelp) I162B(n int16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &n)
	return bytesBuffer.Bytes()
}

func (p *ByteHelp) B2I16(bs []byte) int16 {
	bytesBuffer := bytes.NewBuffer(bs)
	var v int16
	binary.Read(bytesBuffer, binary.BigEndian, &v)
	return v
}

func (p *ByteHelp) I642B(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &n)
	return bytesBuffer.Bytes()
}

func (p *ByteHelp) B2I64(bs []byte) int64 {
	bytesBuffer := bytes.NewBuffer(bs)
	var v int64
	binary.Read(bytesBuffer, binary.BigEndian, &v)
	return v
}

func (p *ByteHelp) UI32ToB(n uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &n)
	return bytesBuffer.Bytes()
}

func (p *ByteHelp) BToUI32(bs []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(bs)
	var v uint32
	binary.Read(bytesBuffer, binary.BigEndian, &v)
	return v
}

func (p *ByteHelp) UI16ToB(n uint16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &n)
	return bytesBuffer.Bytes()
}

func (p *ByteHelp) BToUI16(bs []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(bs)
	var v uint16
	binary.Read(bytesBuffer, binary.BigEndian, &v)
	return v
}

func (p *ByteHelp) BStr(bs []byte) string {
	v := ""
	for _, b := range bs {
		s := fmt.Sprint(uint8(b))
		if len(s) < 2 {
			s = fmt.Sprint("00", s)
		} else if len(s) < 3 {
			s = fmt.Sprint("0", s)
		}
		v = fmt.Sprint(v, " ", s)
	}
	return v
}

func (p *ByteHelp) Append(bss ...[]byte) []byte {
	out := []byte{}
	for _, bs := range bss {
		for _, b := range bs {
			out = append(out, b)
		}
	}
	return out
}

func (p *ByteHelp) ToHex(bs []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range bs {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}
	return buffer.String()
}

func (p *ByteHelp) FromHex(hex string) []byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)
	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}
