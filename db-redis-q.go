package Figo

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisQueue struct {
	rp   *redis.Pool
	name string
}

func NewRedisQueue(rp *redis.Pool, name string, concurrent int, worker func(string, error)) RedisQueue {
	q := RedisQueue{
		rp:   rp,
		name: name,
	}
	f := func() {
		for {
			v, err := q.Deq()
			if err != nil && err != redis.ErrNil {
				worker(v, err)
			} else if v == "" {
				time.Sleep(time.Second)
			} else {
				worker(v, err)
			}
		}
	}
	for i := 0; i < concurrent; i++ {
		go f()
	}
	return q
}

func (p *RedisQueue) Enq(v string) error {
	c := p.rp.Get()
	defer c.Close()
	_, err := c.Do("RPUSH", p.name, v)
	return err
}

func (p *RedisQueue) Deq() (string, error) {
	c := p.rp.Get()
	defer c.Close()
	return redis.String(c.Do("LPOP", p.name))
}
