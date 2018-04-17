package Figo

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"net/url"
)

type HttpHelperMockBuilder struct {
	dataMock map[string]string
}

func NewHttpHelperMockBuilder() HttpHelperMockBuilder {
	return HttpHelperMockBuilder{
		dataMock: make(map[string]string),
	}
}

func (p *HttpHelperMockBuilder) MockVal(key, val string) *HttpHelperMockBuilder {
	p.dataMock[key] = val
	return p
}

func (p *HttpHelperMockBuilder) ParamHelper() ParamHelper {
	m := martini.Params(p.dataMock)
	return ParamHelper{
		param:   m,
		context: make(map[string]string),
	}
}

func (p *HttpHelperMockBuilder) FormHelper() FormHelper {
	values := make(url.Values)
	for k, v := range p.dataMock {
		values[k] = []string{v}
	}
	r := &http.Request{
		Form:     values,
		PostForm: values,
	}
	return FormHelper{
		r:       r,
		context: make(map[string]string),
	}
}

type ResponseWriterMock struct {
}

func (p *ResponseWriterMock) Header() http.Header {
	return make(http.Header)
}

func (p *ResponseWriterMock) Write(b []byte) (int, error) {
	log.Println(string(b))
	return 200, nil
}
func (p *ResponseWriterMock) WriteHeader(code int) {
	log.Println("http header code :", code)
}
