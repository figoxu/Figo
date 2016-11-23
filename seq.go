package Figo

import (
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
	"sync"
)

type Seq interface {
	Next() int64
}

type SeqRedis struct {
	rp  *redis.Pool
	key string
}

func NewSeqRedis(rp *redis.Pool, key string) *SeqRedis {
	return &SeqRedis{
		rp: rp,
		key, key,
	}
}

func (p *SeqRedis) Next() int64 {
	c := p.rp.Get()
	defer c.Close()
	v, err := redis.Int64(c.Do("INCR", p.key))
	utee.Chk(err)
	return v
}

type SeqMem struct {
	lock    sync.Mutex
	counter int64
}

func NewSeqMem() *SeqMem {
	return &SeqMem{}
}

func (p *SeqMem) Next() int64 {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.counter++
	return p.counter
}
