package Figo

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"log"
	"testing"
)

func TestNewBlockExecuteQ(t *testing.T) {

	bq := NewBlockExecuteQ(1000, 3, 3, func(v interface{}) bool {
		b := randomdata.Boolean()
		if b {
			log.Println("execute @v:", v, " SUCCESS")
		} else {
			log.Println("execute @v:", v, " FAILURE")
		}
		return b
	})

	mockInput := func() {
		for i := 0; i < 30; i++ {
			bq.Enq(fmt.Sprint("testData ", i))
		}
	}
	mockInput()
	<-make(chan bool)
}
