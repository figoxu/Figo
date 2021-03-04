package Figo

import (
	"encoding/json"
	"github.com/quexer/utee"
	"io"
	"net/http"
	"net/http/httptest"
)

func MockRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	utee.Chk(err)
	return req
}

func MockUteeWeb() Web {
	return Web{W: httptest.NewRecorder()}
}

type Web struct {
	W http.ResponseWriter
}

func (p *Web) Json(code int, data interface{}) (int, string) {
	b, err := json.Marshal(data)
	utee.Chk(err)
	p.W.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return code, string(b)
}
func (p *Web) Txt(code int, txt string) (int, string) {
	p.W.Header().Set("Content-Type", "html/text; charset=UTF-8")
	return code, txt
}
