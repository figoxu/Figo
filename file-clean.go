package Figo

import (
	"time"
	"os"
	"path/filepath"
	"strings"
	"github.com/quexer/utee"
	"io"
)

type FileCleanHelper struct {
	ScanDir          string
	BackDir          string
	BackDurationSec  int
	CleanDurationSec int
}

func NewFileCleanHelper(ScanDir, BackDir string, BackDurationSec, CleanDurationSec int) *FileCleanHelper {
	return &FileCleanHelper{
		ScanDir:          NewFilePath(ScanDir).UnixPath(),
		BackDir:          NewFilePath(BackDir).UnixPath(),
		BackDurationSec:  BackDurationSec,
		CleanDurationSec: CleanDurationSec,
	}
}

func (p *FileCleanHelper) Back() {
	defer Catch()
	dateLine := time.Now().Add(time.Second * time.Duration(p.BackDurationSec))
	filepath.Walk(p.ScanDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if dateLine.Unix()-info.ModTime().Unix() > 0 {
			outFilePath := NewFilePath(strings.Replace(path,p.ScanDir,p.BackDir,-1))
			newPath := outFilePath.UnixPath()
			err:=os.MkdirAll(outFilePath.FolderName(), 0777)
			utee.Chk(err)
			newFileName := strings.Replace(newPath, p.ScanDir, p.BackDir, -1)
			in,err:=NewFilePath(path).Open()
			utee.Chk(err)
			out,err:=NewFilePath(newFileName).Open()
			utee.Chk(err)
			_, err = io.Copy(out, in)
			utee.Chk(err)
			os.Remove(path)
		}
		return nil
	})
}

func (p *FileCleanHelper) Clean() {
	defer Catch()
	dateLine := time.Now().Add(time.Second * time.Duration(p.CleanDurationSec))
	filepath.Walk(p.BackDir, func(path string, info os.FileInfo, err error) error {
		if dateLine.Unix()-info.ModTime().Unix() > 0 {
			os.Remove(path)
		}
		return nil
	})
}
