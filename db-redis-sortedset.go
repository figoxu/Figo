package Figo

import (
	"github.com/garyburd/redigo/redis"
)

type RedisSortedSet struct {
	key string
	rp  *redis.Pool
}

func NewRedisSortedSet(key string, rp *redis.Pool) *RedisSortedSet {
	return &RedisSortedSet{
		key: key,
		rp:  rp,
	}
}

func (p *RedisSortedSet) ZAdd(score int64, name string) {
	c := p.rp.Get()
	defer c.Close()
	c.Do("ZADD", p.key, score, name)
}

func (p *RedisSortedSet) ZBatchAdd(ssitems ...SSItem) {
	c := p.rp.Get()
	defer c.Close()
	for _, ssitem := range ssitems {
		c.Send("ZADD", p.key, ssitem.Score, ssitem.Key)
	}
	c.Flush()
}

func (p *RedisSortedSet) ZRem(name string) {
	c := p.rp.Get()
	defer c.Close()
	c.Do("ZREM", p.key, name)
}

func (p *RedisSortedSet) ZScore(name string) int64 {
	c := p.rp.Get()
	defer c.Close()
	if score, err := redis.Int64(c.Do("ZSCORE", p.key, name)); err == nil {
		return score
	}
	return 0
}

func (p *RedisSortedSet) ZCount(min, max int64) int64 {
	c := p.rp.Get()
	defer c.Close()
	count := func(key string) int64 {
		if v, err := redis.Int64(c.Do("ZCOUNT", key, min, max)); err == nil {
			return v
		}
		return 0
	}
	total := count(p.key)
	return total
}

func (p *RedisSortedSet) ZRangeByScore(min, max int64) []string {
	c := p.rp.Get()
	defer c.Close()
	if skeys, err := redis.Strings(c.Do("ZRANGEBYSCORE", p.key, min, max)); err == nil {
		return skeys
	}
	return []string{}
}
