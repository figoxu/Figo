package Figo

import (
	"github.com/garyburd/redigo/redis"
)

const SHORT_URL_COUNT_KEY string = "short_url_count"

type Sequence interface {
	Next() int
}

type SequenceObj struct {
}

func (s *SequenceObj) Next(rp *redis.Pool) int64 {
	c := rp.Get()
	defer c.Close()
	v, _ := redis.Int64(c.Do("INCR", SHORT_URL_COUNT_KEY))
	return v
}
