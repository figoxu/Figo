package Figo

import (
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

const (
	PREFIX = "tiny"
)

type tinyUrl struct {
	cache Cache
}

func NewTinyUrl(cache Cache) tinyUrl {
	return tinyUrl{cache: cache}
}

func (t tinyUrl) Convert(url string, rp *redis.Pool) string {
	//uid := uuid.NewUUID().String()
	//k := fmt.Sprintf("%s-%s", PREFIX, uid)
	s := &SequenceObj{}
	k := s.Next(rp)
	t.cache.Put(strconv.FormatInt(k, 16), url)
	return strconv.FormatInt(k, 16)
}

func (t tinyUrl) GetRedirectUrlHandler(key string) martini.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		originUrl := t.getUrl(key)
		http.Redirect(w, r, originUrl, 301)
	}
}

func (t *tinyUrl) getUrl(key string) string {
	if v := t.cache.Get(key); v != nil {
		return v.(string)
	}
	return ""
}
