package Figo

import (
	"github.com/colinmarc/hdfs"
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
	folderName := NewFilePath(fullPath).FolderName()
	p.client.MkdirAll(folderName, 0644)
	w, err := p.client.Create(fullPath)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	w.Close()
	return err
}

func (p *HDFSClient) Read(fullPath string) ([]byte, error) {
	p.open()
	defer p.close()
	return p.client.ReadFile(fullPath)
}
