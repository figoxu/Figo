package Figo

import (
	as "github.com/aerospike/aerospike-client-go"
	"github.com/aerospike/aerospike-client-go/types"
	"github.com/quexer/utee"
	"log"
)

var AsUtee = &AerospikeUtil{}

type AerospikeUtil struct{}
type AsSetInfo struct {
	NameSpace string
	SetName   string
}

func (p *AerospikeUtil) AsConnect(s string) *as.Client {
	h, port, err := ParseUrl(s)
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

func (p *AerospikeUtil) Index(ac *as.Client, setInfo AsSetInfo, binName string, tp as.IndexType) {
	if idxTask, err := ac.CreateIndex(nil, setInfo.NameSpace, setInfo.SetName, setInfo.SetName+binName, binName, tp); err == nil {
		<-idxTask.OnComplete()
	} else {
		if asErr, ok := err.(types.AerospikeError); !ok || asErr.ResultCode() != types.INDEX_FOUND {
			panic(err)
		}
	}
}

func (p *AerospikeUtil) Put(ac *as.Client, setInfo AsSetInfo, key string, val interface{}) error {
	asKey, err := as.NewKey(setInfo.NameSpace, setInfo.SetName, key)
	switch v := val.(type) {
	case as.BinMap:
		err = ac.Put(nil, asKey, v)
	default:
		err = ac.PutObject(nil, asKey, val)
	}
	return err
}

func (p *AerospikeUtil) Get(ac *as.Client, setInfo AsSetInfo, key string, val interface{}) error {
	asKey, err := as.NewKey(setInfo.NameSpace, setInfo.SetName, key)
	ac.GetObject(nil, asKey, val)
	return err
}

func (p *AerospikeUtil) NewSetInfo(namespace, setName string) AsSetInfo {
	return AsSetInfo{
		NameSpace: namespace,
		SetName:   setName,
	}
}
