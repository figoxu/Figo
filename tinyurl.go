package Figo

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/pborman/uuid"
	"net/http"
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

func (t tinyUrl) Convert(url string) string {
	uid := uuid.NewUUID().String()
	k := fmt.Sprintf("%s-%s", PREFIX, uid)
	t.cache.Put(k, url)
	return k
}

func (t tinyUrl) GetRedirectUrlHandler(key string) martini.Handler {
	originUrl := t.getUrl(key)
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, originUrl, 301)
	}
}

func (t *tinyUrl) getUrl(key string) string {
	if v := t.cache.Get(key); v != nil {
		return v.(string)
	}
	return ""
}
