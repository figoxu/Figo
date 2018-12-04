package Figo

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/quexer/utee"
	"net/http"
	"strconv"
	"strings"
	"net/url"
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


func UrlAppendParam(rawUrl, k, v string) string {
	reqURI, err := url.ParseRequestURI(rawUrl)
	utee.Chk(err)
	vs := reqURI.Query()
	vs.Set(k, v)
	reqURI.RawQuery = vs.Encode()
	return reqURI.String()
}

func UrlRemoveParam(rawUrl,k string)string{
	pureUrl,hash := rawUrl,""
	if idx:=strings.LastIndex(rawUrl,"#");idx!=-1 {
		pureUrl = rawUrl[0:idx]
		hash = rawUrl[idx:len(rawUrl)]
	}
	reqURI, err := url.ParseRequestURI(pureUrl)
	utee.Chk(err)
	vs := reqURI.Query()
	vs.Del(k)
	reqURI.RawQuery = vs.Encode()
	return fmt.Sprint(reqURI.String(),hash)
}

func UrlExistParam(rawUrl,k string)bool{
	reqURI, err := url.ParseRequestURI(rawUrl)
	utee.Chk(err)
	vs := reqURI.Query()
	return vs.Get(k)!=""
}
