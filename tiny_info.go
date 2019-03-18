package Figo

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/red"
)

type RedTinyInfo struct {
	redDo red.DoFunc
	key   string
}

func NewTinyInfo(redDo red.DoFunc, key string) *RedTinyInfo {
	return &RedTinyInfo{
		redDo: redDo,
		key:   key,
	}
}

func (p *RedTinyInfo) keyStore() string {
	return fmt.Sprint("red_tiny_", p.key, "_store")
}

func (p *RedTinyInfo) keySeq() string {
	return fmt.Sprint("red_tiny_", p.key, "_seq")
}

func (p *RedTinyInfo) Put(content string) (int, error) {
	v, err := redis.Int(p.redDo("INCR", p.keySeq()))
	if err != nil {
		return v, err
	}
	_, err = p.redDo("HSET", p.keyStore(), v, content)
	return v, err
}

func (p *RedTinyInfo) Get(seq int) (string, error) {
	return redis.String(p.redDo("HGET", p.keyStore(), seq))
}
