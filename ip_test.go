package Figo

import (
	"testing"
	"log"
)

func TestIpHelp_FromBytes(t *testing.T) {
	bs := Ip.ToBytes("8.8.8.8:65535")
	log.Println(bs, "   ", Ip.FromBytes(bs))
	bs = Ip.ToBytes("255.255.255.0:255")
	log.Println(bs, "   ", Ip.FromBytes(bs))
	bs = Ip.ToBytes("127.0.0.1:256")
	log.Println(bs, "   ", Ip.FromBytes(bs))
	bs = Ip.ToBytes("192.168.0.1")
	log.Println(bs, "   ", Ip.FromBytes(bs))
}

func TestIpHelp_ToBytes(t *testing.T) {
	log.Println(Ip.ToBytes("8.8.8.8:65535"))
	log.Println(Ip.ToBytes("255.255.255.0:255"))
	log.Println(Ip.ToBytes("127.0.0.1:256"))
	log.Println(Ip.ToBytes("192.168.0.1"))
}
