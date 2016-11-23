package Figo

import (
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
)

type Seq interface {
	Next() int64
}

type SeqRedis struct {
	rp  *redis.Pool
	key string
}

func (p *SeqRedis) Next() int64 {
	c := p.rp.Get()
	defer c.Close()
	v, err := redis.Int64(c.Do("INCR", p.key))
	utee.Chk(err)
	return v
}
