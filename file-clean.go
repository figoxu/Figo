package Figo

import (
	"time"
	"os"
	"path/filepath"
	"strings"
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
	dateLine := time.Now().Add(time.Second * time.Duration(p.BackDurationSec))
	filepath.Walk(p.ScanDir, func(path string, info os.FileInfo, err error) error {
		if dateLine.Unix()-info.ModTime().Unix() > 0 {
			newPath := NewFilePath(path).UnixPath()
			newFileName := strings.Replace(newPath, p.ScanDir, p.BackDir, -1)
			os.Rename(path, newFileName)
		}
		return nil
	})
}

func (p *FileCleanHelper) Clean() {
	dateLine := time.Now().Add(time.Second * time.Duration(p.CleanDurationSec))
	filepath.Walk(p.ScanDir, func(path string, info os.FileInfo, err error) error {
		if dateLine.Unix()-info.ModTime().Unix() > 0 {
			os.Remove(path)
		}
		return nil
	})
}
