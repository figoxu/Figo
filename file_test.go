package Figo

import (
	"github.com/quexer/utee"
	"log"
	"testing"
)

func TestFileOpen(t *testing.T) {
	_, err := FileOpen("./test.txt")
	utee.Chk(err)
	fp := FilePath("./test.txt")
	fullPath, err := fp.FullPath()
	utee.Chk(err)
	log.Println(fullPath)
	fp = FilePath(fullPath)
	log.Println(fp.WindowsPath())
	log.Println(fp.UnixPath())
	fp = FilePath(fp.UnixPath())
	log.Println(fp.WindowsPath())
}
