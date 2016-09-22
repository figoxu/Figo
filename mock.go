package Figo

import (
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

func MockUteeWeb() utee.Web {
	return utee.Web{W: httptest.NewRecorder()}
}
