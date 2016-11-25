package Figo

import (
	"fmt"
	"github.com/figoxu/utee"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
	"strings"
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

func (p *TinyUrl) Handler() martini.Handler {
	return func(w http.ResponseWriter, r *http.Request, param martini.Params) {
		defer Catch()
		if val := p.cache.Get(param["key"]); val != nil {
			originUrl, err := TpString(val)
			utee.Chk(err)
			if strings.Index(originUrl, "http") == -1 {
				originUrl = fmt.Sprint("http://", originUrl)
			}
			http.Redirect(w, r, originUrl, 301)
		}
	}
}
