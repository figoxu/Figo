package Figo
import (
	"log"
	"testing"
	"github.com/quexer/utee"
)

func TestTpInt(t *testing.T) {
	v,err :=TpInt(int64(1000))
	utee.Chk(err)
	log.Println(v==1000)
	v,err =TpInt("1024")
	utee.Chk(err)
	log.Println(v==1024)
	v,err =TpInt([]byte("768"))
	utee.Chk(err)
	log.Println(v==768)
}
