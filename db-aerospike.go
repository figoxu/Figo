package Figo

import (
	as "github.com/aerospike/aerospike-client-go"
	"github.com/aerospike/aerospike-client-go/types"
	"github.com/quexer/utee"
	"log"
)

var AS = &AerospikeUtil{}

type AerospikeUtil struct{}

func (p *AerospikeUtil) AsConnect(s string) *as.Client {
	h, port, err := utee.ParseUrl(s)
	utee.Chk(err)
	ac, err := as.NewClient(h, port)
	utee.Chk(err)
	log.Println("[as nodes]:", ac.GetNodes())
	return ac
}

func (p *AerospikeUtil) InitLua(ac *as.Client, serverPath, udfCode string) {
	regTask, err := ac.RegisterUDF(nil, []byte(udfCode), serverPath, as.LUA)
	utee.Chk(err)
	err = <-regTask.OnComplete()
	utee.Chk(err)
}

func (p *AerospikeUtil) Index(ac *as.Client, namespace, setName, binName string, tp as.IndexType) {
	if idxTask, err := ac.CreateIndex(nil, namespace, setName, setName+binName, binName, tp); err == nil {
		<-idxTask.OnComplete()
	} else {
		if asErr, ok := err.(types.AerospikeError); !ok || asErr.ResultCode() != types.INDEX_FOUND {
			panic(err)
		}
	}
}
