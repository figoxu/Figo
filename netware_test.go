package Figo

import "testing"

func TestTcpProxy(t *testing.T) {
	localAddr := "127.0.0.1:8080"
	remoteAddr := "127.0.0.1:6379"
	redisProxy := NewTcpProxy(localAddr, remoteAddr, 100)
	redisProxy.Listen()
}
