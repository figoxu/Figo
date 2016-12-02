package Figo

import "testing"

func TestSMTPClient(t *testing.T) {
	client := GetSMTPClient("xujh945@qq.com", "xxxxxxx", "smtp.qq.com:25")
	client.Send("xxx@qq.com", "TestMsg", "Hello World")
}
