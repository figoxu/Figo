package Figo

import (
	"fmt"
	"strings"
)

var Ip = IpHelp{}

type IpHelp struct {
}

func (p *IpHelp) ToBytes(addr string) (bs []byte) {
	ips, port := p.ToHostPort(addr)
	h, l := byte(port>>8), byte(port&0xff)
	bs = make([]byte, 0)
	bs = append(bs, ips...)
	bs = append(bs, h, l)
	return bs
}

func (p *IpHelp) ToHostPort(addr string) (ips []byte, port uint16) {
	vs := strings.Split(addr, ":")
	host := vs[0]
	port = 80
	if len(vs) > 1 {
		portV, _ := TpInt(vs[1])
		port = uint16(portV)
	}
	ipSegments := strings.Split(host, ".")
	ips = make([]byte, 0)
	for _, ipStr := range ipSegments {
		v, _ := TpInt(ipStr)
		ips = append(ips, byte(v))
	}
	return ips, port
}

func (p *IpHelp) FromBytes(bs []byte) (addr string) {
	addr = fmt.Sprint(bs[0], ".", bs[1], ".", bs[2], ".", bs[3])
	if len(bs) == 6 {
		addr = fmt.Sprint(addr, ":", int(bs[4])*256+int(bs[5]))
	} else {
		addr = fmt.Sprint(addr, ":80")
	}
	return addr
}
