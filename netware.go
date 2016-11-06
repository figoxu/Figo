package Figo

import (
	"github.com/quexer/utee"
	"io"
	"net"
)

type TcpProxy struct {
	local      net.Listener
	remoteAddr string
	latch      *utee.Throttle
}

func NewTcpProxy(localAddr, remoteAddr string, concurrent int) *TcpProxy {
	local, err := net.Listen("tcp", localAddr)
	utee.Chk(err)
	latch := utee.NewThrottle(concurrent)
	return &TcpProxy{
		local:      local,
		remoteAddr: remoteAddr,
		latch:      latch,
	}
}

func (p *TcpProxy) Listen() {
	for {
		conn, err := p.local.Accept()
		if conn == nil {
			utee.Chk(err)
		}
		if p.latch.Available() <= 0 {
			conn.Write([]byte("Server Is Busy"))
			conn.Close()
			continue
		}
		go p.forward(conn)
	}
}

func (p *TcpProxy) forward(local net.Conn) {
	defer Catch()
	p.latch.Acquire()
	defer p.latch.Release()
	remote, err := net.Dial("tcp", p.remoteAddr)
	utee.Chk(err)
	go io.Copy(local, remote)
	io.Copy(remote, local)
}
