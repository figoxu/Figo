package Figo

import "github.com/garyburd/redigo/redis"

type RedisMutex struct {
	rp       *redis.Pool
	resource string
}

func (p *RedisMutex) Lock(ttlSec int) (bool, error) {
	c := p.rp.Get()
	defer c.Close()
	if _, err := redis.String(c.Do("SET", p.resource, "1", "EX", ttlSec, "NX")); err != nil {
		if err == redis.ErrNil {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *RedisMutex) Unlock() {
	c := p.rp.Get()
	defer c.Close()
	c.Do("DEL", p.resource)
}
