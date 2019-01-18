package Figo

import (
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

// 支持'去重'的优先级队列
type RedisZQueue struct {
	rp   *redis.Pool
	name string
}

var zpopScript = redis.NewScript(1, `
    local r = redis.call('ZRANGE', KEYS[1], 0, 0)
    if r ~= nil then
        r = r[1]
        redis.call('ZREM', KEYS[1], r)
    end
    return r
`)

func NewRedisZQueue(rp *redis.Pool, name string, concurrent int, worker func(string, error)) RedisZQueue {
	q := RedisZQueue{
		rp:   rp,
		name: name,
	}
	f := func() {
		for {
			v, err := q.Deq()
			if v == "" {
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

func (p *RedisZQueue) Enq(v string, score int) error {
	c := p.rp.Get()
	defer c.Close()
	_, err := c.Do("ZADD", p.name, score, v)
	return err
}

func (p *RedisZQueue) Deq() (string, error) {
	c := p.rp.Get()
	defer c.Close()
	result, err := redis.String(zpopScript.Do(c, p.name))
	if err == redis.ErrNil {
		return "", nil
	} else if err != nil {
		errCauseEmptyFlag := strings.Index(err.Error(), "arguments must be strings or integers") != -1
		if errCauseEmptyFlag {
			return "", nil
		}
	}
	return result, err
}
