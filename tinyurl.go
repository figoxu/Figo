package Figo

import (
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

const (
	PREFIX = "tiny"
)

type TinyUrl struct {
	cache Cache
	seq   Seq
}

func NewTinyUrl(cache Cache, seq Seq) TinyUrl {
	return TinyUrl{cache: cache, seq: seq}
}

func (p *TinyUrl) Convert(url string) string {
	k := p.seq.Next()
	key := strconv.FormatInt(k, 16)
	p.cache.Put(key, url)
	return key
}

func (p *TinyUrl) Handler(key string) martini.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if val := p.cache.Get(key); val != nil {
			originUrl := val.(string)
			http.Redirect(w, r, originUrl, 301)
		}
	}
}
