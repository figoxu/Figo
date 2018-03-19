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
		Float64: wp_func_float64(m),
		Bool:    wp_func_Bool(m),
		Int:     wp_func_Int(m),
		Time:    wp_func_time(m),
		TimeLoc: wp_func_time_loc(m),
		String:  wp_func_string(m),
		IntArr:  wp_func_IntArray(m),
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
		Int:     form_func_Int(r),
		Float32: form_func_Float32(r),
		String:  form_func_String(r),
		StrArr:  form_func_StrArray(r),
		IntArr:  form_func_IntArray(r),
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
