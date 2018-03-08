package Figo

import (
	"testing"
	"github.com/quexer/utee"
	"log"
)

func TestFileUtee_WriteLinesSlice(t *testing.T) {
	futee:=FileUtee{}
	path := "/server/nfs/session/test.txt"
	futee.WriteLinesSlice([]string{"hello","world","how","r","u"}, path)
	futee.WriteLinesSlice([]string{"figo","xu","is","awesome"}, path)
	txts,err:=futee.ReadLinesSlice(path)
	utee.Chk(err)
	for _,txt:=range txts{
		log.Println(txt)
	}
}