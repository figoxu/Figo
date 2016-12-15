package Figo

import (
	"github.com/garyburd/redigo/redis"
)

type IdService struct {
	cache Cache
	seq   Seq
}

const OFFSET_NOT_FOUND = -1

func NewIdService(cache Cache, seq Seq) *IdService {
	return &IdService{
		cache: cache,
		seq:   seq,
	}
}

func (p *IdService) GetOffSet(key string) int64 {
	if v := p.cache.Get(key); v != nil && v != OFFSET_NOT_FOUND {
		return v.(int64)
	} else {
		offset := p.seq.Next()
		p.cache.Put(key, offset)
		return offset
	}
}

type RedisBitMap struct {
	rp  *redis.Pool
	key string
}

func NewRedisBitMap(rp *redis.Pool, key string) *RedisBitMap {
	return &RedisBitMap{
		rp:  rp,
		key: key,
	}
}

func (p *RedisBitMap) Clear() error {
	c := p.rp.Get()
	defer c.Close()
	_, err := c.Do("DEL", p.key)
	return err
}

func (p *RedisBitMap) Set(offset int, val bool) error {
	c := p.rp.Get()
	defer c.Close()
	if val {
		return c.Send("setbit", p.key, offset, 1)
	}
	return c.Send("setbit", p.key, offset, 0)
}

func (p *RedisBitMap) Get(offset int) int {
	c := p.rp.Get()
	defer c.Close()
	v, _ := redis.Int(c.Do("getbit", p.key, offset))
	return v
}

func (p *RedisBitMap) Count() int {
	c := p.rp.Get()
	defer c.Close()
	count, _ := redis.Int(c.Do("bitcount", p.key))
	return count
}

func (p *RedisBitMap) And(resultKey string, keys ...string) *RedisBitMap {
	c := p.rp.Get()
	defer c.Close()
	params := []interface{}{"and", resultKey, p.key}
	for _, v := range keys {
		params = append(params, v)
	}
	c.Send("bitop", params...)
	return &RedisBitMap{
		rp:  p.rp,
		key: resultKey,
	}
}

func (p *RedisBitMap) Or(resultKey string, keys ...string) *RedisBitMap {
	c := p.rp.Get()
	defer c.Close()
	params := []interface{}{"or", resultKey, p.key}
	for _, v := range keys {
		params = append(params, v)
	}
	c.Send("bitop", params...)
	return &RedisBitMap{
		rp:  p.rp,
		key: resultKey,
	}
}
