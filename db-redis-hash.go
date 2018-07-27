package Figo

import (
	"github.com/garyburd/redigo/redis"
	"github.com/murlokswarm/errors"
)

type RedisHash struct {
	rp        *redis.Pool
	masterKey string
}

type Skv struct {
	K string
	V string
}

func NewRedisHash(rp *redis.Pool, masterKey string) *RedisHash {
	return &RedisHash{
		rp:        rp,
		masterKey: masterKey,
	}
}

func (p *RedisHash) Set(kvs ...Skv) (interface{}, error) {
	c := p.rp.Get()
	defer c.Close()
	if len(kvs) <= 0 {
		return "", errors.New("Bad Param For Redis Hash As Set")
	} else if len(kvs) == 1 {
		kv := kvs[0]
		return c.Do("HSET", p.masterKey, kv.K, kv.V)
	}
	args := make([]interface{}, 0)
	args = append(args, p.masterKey)
	for _, kv := range kvs {
		args = append(args, kv.K, kv.V)
	}
	return c.Do("HMSET", args...)
}

func (p *RedisHash) Get(ks ...string) ([]Skv, error) {
	c := p.rp.Get()
	defer c.Close()
	kvs := make([]Skv, 0)
	if len(ks) <= 0 {
		return kvs, errors.New("Bad Param For Redis Hash As Get")
	} else if len(ks) == 1 {
		k := ks[0]
		v, err := redis.String(c.Do("HGET", p.masterKey, k))
		kvs = append(kvs, Skv{
			K: k,
			V: v,
		})
		return kvs, err
	}

	args := make([]interface{}, 0)
	args = append(args, p.masterKey)
	for _, k := range ks {
		args = append(args, k)
	}
	vs, err := redis.Strings(c.Do("HMGET", args...))
	if len(ks) != len(vs) || err != nil {
		return kvs, err
	}
	for i, k := range ks {
		kvs = append(kvs, Skv{
			K: k,
			V: vs[i],
		})
	}
	return kvs, nil
}
