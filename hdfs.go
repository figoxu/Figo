package Figo

import (
	"github.com/colinmarc/hdfs"
	"io"
	"log"
	"os"
)

type HDFSClient struct {
	client   *hdfs.Client
	nameNode string
	user     string
}

func NewHDFSClient(nameNode, user string) HDFSClient {
	return HDFSClient{
		nameNode: nameNode,
		user:     user,
	}
}

func (p *HDFSClient) open() error {
	if p.client != nil {
		return nil
	}
	os.Setenv("HADOOP_USER_NAME", p.user)
	client, err := hdfs.New(p.nameNode)
	p.client = client
	return err
}

func (p *HDFSClient) close() error {
	if p.client != nil {
		err := p.client.Close()
		p.client = nil
		return err
	}
	return nil
}

func (p *HDFSClient) Write(fullPath string, data []byte) error {
	p.open()
	defer p.close()
	fullPath = FilePathFormat(fullPath)
	folderName := NewFilePath(fullPath).FolderName()
	if err := p.client.MkdirAll(folderName, 0644); err != nil {
		return err
	}
	w, err := p.client.Create(fullPath)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	w.Close()
	return err
}

func (p *HDFSClient) AppendFile(fullPath string, data []byte) error {
	p.open()
	defer p.close()
	fullPath = FilePathFormat(fullPath)
	w, err := p.client.Append(fullPath)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = w.Write(data)
	return err
}

func (p *HDFSClient) WriteFile(fullPath string, file *os.File) error {
	p.open()
	defer p.close()
	fullPath = FilePathFormat(fullPath)
	folderName := NewFilePath(fullPath).FolderName()
	if err := p.client.MkdirAll(folderName, 0644); err != nil {
		return err
	}
	w, err := p.client.Create(fullPath)
	defer w.Close()
	if err != nil {
		return err
	}
	n, err := io.Copy(w, file)
	log.Println("@WriteFile ", n, " Bytes")
	return err
}

func (p *HDFSClient) Read(fullPath string) ([]byte, error) {
	p.open()
	defer p.close()
	return p.client.ReadFile(fullPath)
}
