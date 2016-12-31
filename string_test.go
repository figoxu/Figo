package Figo

import (
	"log"
	"testing"
)

func TestSnakeString(t *testing.T) {
	log.Println(SnakeString("FigoHowAreYou"))
}

func TestCamelString(t *testing.T) {
	log.Println(CamelString("FigoHowAreYou"))
}
