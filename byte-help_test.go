package Figo

import (
	"log"
	"testing"
	"time"
)

func TestByteHelp_Append(t *testing.T) {
	bs := Bh.Append([]byte{1, 2, 3, 4, 5}, []byte{6, 7, 8}, []byte{9, 10, 11, 12, 13})
	log.Println(bs)
	str := Bh.BStr(bs)
	log.Println(str)
}

func TestByteHelp_B2I16(t *testing.T) {
	v := Bh.B2I16([]byte{1, 22})
	log.Println(v)
	v = Bh.B2I16([]byte{1, 0})
	log.Println(v)
}

func TestByteHelp_I162B(t *testing.T) {
	v := Bh.I162B(int16(278))
	log.Println(Bh.BStr(v))
	v = Bh.I162B(int16(256))
	log.Println(Bh.BStr(v))
}

func TestByteHelp_ToHex(t *testing.T) {
	bs := Bh.Append([]byte{1})
	v := Bh.ToHex(bs)
	bs2 := Bh.FromHex(v)
	log.Println("@v:", v, " @result:", bs2)

	for i := 0; i < 20; i++ {
		bs = Bh.Append(bs, []byte{byte(i)})
		v := Bh.ToHex(bs)
		bs2 := Bh.FromHex(v)
		log.Println("@v:", v, " @result:", bs2)
	}
}

func TestByteHelp_BToUI32(t *testing.T) {
	v := uint32(102345)
	bs := Bh.UI32ToB(v)
	v2 := Bh.BToUI32(bs)
	log.Println(v, "  to bytes :", bs, " to uint32 :", v2)
}

func TestByteHelp_I642B(t *testing.T) {
	v := time.Now().Unix()
	bs := Bh.I642B(v)
	log.Println(v)
	log.Println(bs)
	log.Println(Bh.B2I64(bs))
}
