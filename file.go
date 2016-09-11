package Figo

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type FilePath string

func (p FilePath) FullPath() (string, error) {
	f, err := p.Open()
	if err != nil {
		return "", err
	}
	return filepath.Abs(f.Name())
}

func (p FilePath) UnixPath() string {
	return strings.Replace(p.String(), "\\", "/", -1)
}

func (p FilePath) WindowsPath() string {
	return strings.Replace(p.String(), "/", "\\", -1)
}

func (p FilePath) FileName() string {
	toks := strings.Split(p.UnixPath(), "/")
	return toks[len(toks)-1]
}

func (p FilePath) FolderName() string {
	return strings.Replace(p.String(), p.FileName(), "", -1)
}

func (p FilePath) Exist() bool {
	var exist = true
	if _, err := os.Stat(p.String()); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func (p FilePath) String() string {
	return string(p)
}

func (p FilePath) Writer() (*bufio.Writer, error) {
	if f, err := p.Open(); err != nil {
		return nil, err
	} else {
		return bufio.NewWriter(f), nil
	}
}

func (p FilePath) Open() (*os.File, error) {
	var file *os.File
	var err error
	if p.Exist() {
		file, err = os.OpenFile(p.String(), os.O_RDWR, 0666)
	} else {
		if err := os.MkdirAll(p.FolderName(), 0777); err != nil {
			return nil, err
		}
		file, err = os.Create(p.String()) //创建文件
	}
	return file, err
}

func FileOpen(s string) (*os.File, error) {
	var filepath FilePath = FilePath(s)
	return filepath.Open()
}

func NewFilePath(s string) FilePath {
	return FilePath(s)
}
