package Figo

import (
	"testing"
	"log"
	"time"
)

func TestNewBothWayMap(t *testing.T) {
	bothwayMap :=NewBothWayMap(0)
	bothwayMap.Put("Hello","world")
	log.Println(bothwayMap.GetByKey("Hello"))
	log.Println(bothwayMap.GetByValue("world"))
	bothwayMap =NewBothWayMap(3)
	bothwayMap.Put("Hello","world")
	log.Println(bothwayMap.GetByKey("Hello"))
	log.Println(bothwayMap.GetByValue("world"))
	time.Sleep(time.Second*time.Duration(4))
	log.Println(bothwayMap.GetByKey("Hello"))
	log.Println(bothwayMap.GetByValue("world"))
}
