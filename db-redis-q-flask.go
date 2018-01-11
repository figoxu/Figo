package Figo

import (
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
)

type RedisQFlask struct {
	rp        *redis.Pool
	name      string
	flaskSize int
}

func NewRedisQFlask(rp *redis.Pool, name string, flaskSize int) RedisQFlask {
	return RedisQFlask{
		rp:        rp,
		name:      name,
		flaskSize: flaskSize,
	}
}

func (p *RedisQFlask) Put(v string)  {
	RedisRpush(p.rp,p.name,v)
	for len,_:=RedisLlen(p.rp,p.name);len > p.flaskSize;len,_=RedisLlen(p.rp,p.name) {
		RedisLpop(p.rp,p.name)
	}
}

func (p *RedisQFlask) Get() []string {
	results := make([]string,0)
	l,err:=RedisLlen(p.rp, p.name)
	utee.Chk(err)
	for i:=0;i<l;i++ {
		v,_:=RedisLindex(p.rp,p.name,i)
		results = append(results,v)
	}
	return results
}

type RedisQMultiFlask struct {
	rp        *redis.Pool
	flaskSize int
}

func NewRedisQMultiFlask(rp *redis.Pool,flaskSize int)RedisQMultiFlask{
	return RedisQMultiFlask{
		rp:rp,
		flaskSize:flaskSize,
	}
}

func (p *RedisQMultiFlask) Put(k,v string)  {
	RedisRpush(p.rp,k,v)
	for len,_:=RedisLlen(p.rp,k);len > p.flaskSize;len,_=RedisLlen(p.rp,k) {
		RedisLpop(p.rp,k)
	}
}

func (p *RedisQMultiFlask) Get(k string) []string {
	results := make([]string,0)
	l,err:=RedisLlen(p.rp, k)
	utee.Chk(err)
	for i:=0;i<l;i++ {
		v,err:=RedisLindex(p.rp,k,i)
		utee.Chk(err)
		results = append(results,v)
	}
	return results
}
