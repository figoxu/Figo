package Figo

import (
	"github.com/garyburd/redigo/redis"
)

type RedisMutex struct {
	rp       *redis.Pool
	resource string
}

func NewRedisMutex(rp *redis.Pool, resource string) *RedisMutex {
	return &RedisMutex{
		rp:       rp,
		resource: resource,
	}
}

func (p *RedisMutex) Lock(ttlSec int) (bool, error) {
	lockAction := func() (bool, error) {
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
	lock := func() (bool, error) {
		for {
			if b, err := lockAction(); b == true {
				return true, nil
			} else if err != nil {
				return false, nil
			}
		}
	}
	return lock()
}

func (p *RedisMutex) Unlock() {
	c := p.rp.Get()
	defer c.Close()
	c.Do("DEL", p.resource)
}
