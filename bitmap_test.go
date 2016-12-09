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
