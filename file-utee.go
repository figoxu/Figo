package Figo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"io"
	"github.com/quexer/utee"
	"io/ioutil"
)

type FileUtee struct {
}

func (p *FileUtee) Exist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func (p *FileUtee) MakeFile(dir, fileName string) (*os.File, error) {
	var f *os.File
	var err error
	fileFullName := fmt.Sprint(dir, fileName)
	if p.Exist(fileFullName) { //如果文件存在
		f, err = os.OpenFile(fileFullName, os.O_RDWR, 0666) //打开文件
		log.Println("@fileFullPath:", fileFullName, " is exist")
	} else {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return nil, err
		}
		f, err = os.Create(fileFullName) //创建文件
	}
	return f, err
}

func (p *FileUtee) ReadLinesChannel(filePath string) (<-chan string, error) {
	c := make(chan string)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	go func() {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()
	return c, nil
}

func (p *FileUtee) ReadLinesSlice(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (p *FileUtee) WriteLinesSlice(lines []string, path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func (p *FileUtee) ReadAll(path string)string{
	fi,err := os.Open(path)
	utee.Chk(err)
	defer fi.Close()
	chunks := make([]byte,1024,1024)
	buf := make([]byte,1024)
	for{
		n,err := fi.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {break}
		chunks=append(chunks,buf[:n]...)
	}
	return string(chunks)
}

func (p *FileUtee) FlushWrite(path,content string)int{
	_,err := os.OpenFile(path,os.O_TRUNC,0666)
	if err!=nil{
		_, err = os.Create(path)
		utee.Chk(err)
	}
	err = ioutil.WriteFile(path, []byte(content),0666)
	utee.Chk(err)
	return len([]byte(content))
}

func (p *FileUtee) FlushWriteBytes(path string,b []byte)int{
	_,err := os.OpenFile(path,os.O_TRUNC,0666)
	if err!=nil{
		_, err = os.Create(path)
		utee.Chk(err)
	}
	err = ioutil.WriteFile(path,b,0666)
	utee.Chk(err)
	return len(b)
}