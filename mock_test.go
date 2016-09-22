package Figo

import (
	"log"
	"net/url"
	"strings"
	"testing"
)

func TestMockRequest(t *testing.T) {
	req := mockRequest("PUT", "http://localhost:8080", strings.NewReader(`foo`))
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	req.Form = url.Values{
		"Hello": []string{"World"},
		"Foo":   []string{"Bar"},
	}
	v := req.FormValue("Hello")
	log.Println(v)
}

func TestMockUteeWeb(t *testing.T) {
	web := mockUteeWeb()
	web.Json(200, []string{"Hello"})
}
