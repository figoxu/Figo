package Figo

import (
	"github.com/quexer/utee"
	"log"
	"testing"
)

func TestHDFSWriteFile(t *testing.T) {
	defer Catch()
	hdfsClient := NewHDFSClient("172.17.0.3:9000", "root")
	fullPath := "/figo/github/bigFile.txt"
	f, e := FileOpen("/home/figo/nohup.out")
	utee.Chk(e)
	err := hdfsClient.WriteFile(fullPath, f)
	utee.Chk(err)
}

func TestHDFSWrite(t *testing.T) {
	hdfsClient := NewHDFSClient("192.168.108.131:9000", "root")
	fullPath := "/figo/github/foo.txt"
	hdfsClient.Write(fullPath, []byte("Here We Go,Frightting To Eat"))
	v, e := hdfsClient.Read(fullPath)
	utee.Chk(e)
	log.Println("@value I Read is :", string(v))
}

func TestHDFSCreateFile(t *testing.T) {
	hdfsClient := NewHDFSClient("", "")
	fullPath := "/test/test/push.txt"
	err := hdfsClient.CreateFile(fullPath)
	utee.Chk(err)
}

func TestHDFSAppend(t *testing.T) {
	hdfsClient := NewHDFSClient("", "")
	fullPath := "/test7.txt"
	err := hdfsClient.AppendFile(fullPath, []byte("append file\n"))
	utee.Chk(err)
}
