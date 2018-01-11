package Figo

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

func RedisPool(server, auth string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     500,
		MaxActive:   500,
		Wait:        true,
		IdleTimeout: 4 * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func RedisGet(rp *redis.Pool, key interface{}) (interface{}, error) {
	c := rp.Get()
	defer c.Close()
	return c.Do("GET", key)
}

func RedisSet(rp *redis.Pool, key, val interface{}) (interface{}, error) {
	c := rp.Get()
	defer c.Close()
	return c.Do("SET", key, val)
}

func RedisSetEx(rp *redis.Pool, key, val interface{}, ttlSec int) (interface{}, error) {
	c := rp.Get()
	defer c.Close()
	return c.Do("SETEX", key, ttlSec, val)
}

func RedisGetbit(rp *redis.Pool, key interface{}, offset int) (bool, error) {
	c := rp.Get()
	defer c.Close()
	v, err := redis.Int(c.Do("getbit", key, offset))
	if err != nil {
		return false, err
	}
	return v == 1, nil
}

func RedisSetbit(rp *redis.Pool, key interface{}, offset int) error {
	c := rp.Get()
	defer c.Close()
	return c.Send("setbit", key, offset, 1)
}

func RedisBitcount(rp *redis.Pool, key interface{}) (int, error) {
	c := rp.Get()
	defer c.Close()
	return redis.Int(c.Do("bitcount", key))
}

func RedisBitop(rp *redis.Pool, tp string, destkey interface{}, key ...interface{}) error {
	c := rp.Get()
	defer c.Close()
	lowerTp := strings.ToLower(tp)
	if !Exist(lowerTp, "or", "and") {
		return errors.New("bad tp")
	}
	keys := []interface{}{lowerTp, destkey}
	for _, v := range key {
		keys = append(keys, v)
	}
	return c.Send("bitop", keys...)
}


func RedisRpush(rp *redis.Pool, k, v string) error {
	c := rp.Get()
	defer c.Close()
	_, err := c.Do("RPUSH", k, v)
	return err
}

func RedisLpop(rp *redis.Pool, k string) (v string,err error) {
	c := rp.Get()
	defer c.Close()
	return redis.String(c.Do("LPOP", k))
}

func RedisLlen(rp *redis.Pool, k string) (int,error) {
	c := rp.Get()
	defer c.Close()

	return redis.Int(c.Do("LLEN",k))
}

func RedisLindex(rp *redis.Pool, k string,index int) (v string,err error){
	c := rp.Get()
	defer c.Close()
	return redis.String(c.Do("LINDEX",k,index))
}