package Figo

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisShardSortedSet struct {
	key   string
	piece int
	rp    *redis.Pool
}

func NewRedisShardSortedSet(key string, piece int, rp *redis.Pool) *RedisShardSortedSet {
	return &RedisShardSortedSet{
		key:   key,
		piece: piece,
		rp:    rp,
	}
}

func (p *RedisShardSortedSet) realKey(name string) string {
	return fmt.Sprint(p.key, "_", Md5ShardPiece(name, p.piece))
}

func (p *RedisShardSortedSet) ZAdd(score int, name string) {
	c := p.rp.Get()
	defer c.Close()
	c.Do("ZADD", p.realKey(name), score, name)
}

type SSItem struct {
	Key   string
	Score int
}

func (p *RedisShardSortedSet) ZBatchAdd(ssitems ...SSItem) {
	c := p.rp.Get()
	defer c.Close()
	for _, ssitem := range ssitems {
		c.Send("ZADD", p.realKey(ssitem.Key), ssitem.Score, ssitem.Key)
	}
	c.Flush()
}

func (p *RedisShardSortedSet) ZScore(name string) int {
	c := p.rp.Get()
	defer c.Close()
	if score, err := redis.Int(c.Do("ZSCORE", p.realKey(name), name)); err == nil {
		return score
	}
	return 0
}

func (p *RedisShardSortedSet) ZCount(min, max int) int {
	c := p.rp.Get()
	defer c.Close()
	count := func(key string) int {
		if v, err := redis.Int(c.Do("ZCOUNT", key, min, max)); err == nil {
			return v
		}
		return 0
	}
	total := 0
	for i := 0; i < p.piece; i++ {
		total += count(fmt.Sprint(p.key, "_", i))
	}
	return total
}
