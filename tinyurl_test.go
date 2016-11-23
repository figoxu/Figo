package Figo

import (
	"flag"
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"github.com/quexer/utee"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	redis_server := "localhost:6379"
	redis_pass := ""
	pool := createPool(redis_server, redis_pass)
	tc := utee.NewTimerCache(3600, nil)
	put := func(key, val interface{}) {
		tc.Put(key, val)
	}
	get := func(key interface{}) interface{} {
		return tc.Get(key)
	}
	cacheObj := NewCacheObj(put, get)
	tinyURL := NewTinyUrl(cacheObj)
	k := tinyURL.Convert("http://www.baidu.com", pool)
	addr := flag.String("p", ":9000", "address where the server listen on")
	flag.Parse()
	m := martini.Classic()
	m.Get("/td", tinyURL.GetRedirectUrlHandler(k))

	http.Handle("/", m)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func createPool(server, auth string) *redis.Pool {
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
