package Figo

import (
	"log"
	"testing"
)

func TestNewIdService(t *testing.T) {
	idService := NewIdService(NewTimerCache(0, nil), NewSeqMem(0))
	log.Println(idService.GetOffSet("figo"), "   ", idService.GetOffSet("zj"), "   ", idService.GetOffSet("zmy"))
	log.Println(idService.GetOffSet("figo"))
}

func TestBitMapService(t *testing.T) {
	rp := RedisPool("127.0.0.1:6379", "")
	bitA := NewRedisBitMap(rp, "BIT_A")
	bitA.Set(1, true)
	bitA.Set(3, true)
	bitA.Set(6, true)
	log.Println("bitA.Count()", bitA.Count())
	bitB := NewRedisBitMap(rp, "BIT_B")
	bitB.Set(2, true)
	bitB.Set(4, true)
	bitB.Set(6, true)
	bitB.Set(7, true)
	log.Println("bitB.Count()", bitB.Count())
	log.Println("BIT_C.Count()", bitA.And("BIT_C", bitB.key).Count())
	log.Println("BIT_D.Count()", bitA.Or("BIT_D", bitB.key).Count())

}
