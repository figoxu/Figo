package Figo

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/figoxu/flog"
	"testing"
)

func TestNewMultiBlockExecuteQ(t *testing.T) {

	type QItem struct {
		k string
		i int
	}
	bmq := NewMultiBlockExecuteQ(1000, 3, 3, func(v interface{}) bool {
		item := v.(QItem)
		var lg = flog.GetLog(item.k)
		b := randomdata.Boolean()
		if b {
			lg.Info("execute @v:", item.i, " SUCCESS")
		} else {
			lg.Info("execute @v:", item.i, " FAILURE")
		}
		return b
	})

	mockInput := func(prefix string) {
		for i := 0; i < 30; i++ {
			bmq.Enq(prefix, QItem{
				k: prefix,
				i: i,
			})
		}
	}
	for i := 0; i < 5; i++ {
		prefix := fmt.Sprint("prefx_", i)
		var testLog = flog.GetLog(prefix)
		testLog.SetConsole(false)
		testLog.SetRollingDaily("./", fmt.Sprint(prefix, ".log"))
		testLog.SetLevel(flog.DEBUG)
		go mockInput(prefix)
	}
	<-make(chan bool)
}
