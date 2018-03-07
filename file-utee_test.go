package Figo

import (
	"testing"
	"github.com/quexer/utee"
	"log"
)

func TestFileUtee_WriteLinesSlice(t *testing.T) {
	futee:=FileUtee{}
	futee.WriteLinesSlice([]string{"hello","world","how","r","u"},"./test.txt")
	futee.WriteLinesSlice([]string{"figo","xu","is","awesome"},"./test.txt")
	txts,err:=futee.ReadLinesSlice("./test.txt")
	utee.Chk(err)
	for _,txt:=range txts{
		log.Println(txt)
	}
}