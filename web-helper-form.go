package Figo

import (
	"net/http"
	"github.com/go-martini/martini"
	"strings"
	"github.com/quexer/utee"
	"strconv"
)

type FormHelper struct {
	context map[string]string
	r       *http.Request
}

func (p *FormHelper) Int(name string, defaultVs ...int) int {
	r, pure := form_func_Int(p.r, name)
	p.context[name] = pure

	return r
}
func (p *FormHelper) Int64(name string) int64 {
	r, pure := form_func_Int64(p.r, name)
	p.context[name] = pure
	return r
}

func (p *FormHelper) Float32(name string, defaultVs ...float32) float32 {
	r, pure := form_func_Float32(p.r, name)
	p.context[name] = pure
	return r
}
func (p *FormHelper) String(name string) string {

	r, pure := form_func_String(p.r, name)
	p.context[name] = pure
	return r
}
func (p *FormHelper) StrArr(name, separate string) []string {
	r, pure := form_func_StrArray(p.r, name, separate)
	p.context[name] = pure
	return r
}
func (p *FormHelper) IntArr(name, separate string) []int {
	r, pure := form_func_IntArray(p.r, name, separate)
	p.context[name] = pure
	return r
}

func (p *FormHelper) Params()map[string]string{
	return p.context
}


func Mid_helper_form(c martini.Context, r *http.Request) {
	c.Map(&FormHelper{
		r:       r,
		context: make(map[string]string),
	})
}



func form_func_Int(r *http.Request, name string, defaultVs ...int) (result int, pure string) {
	pure = r.PostFormValue(name)
	sv := strings.TrimSpace(pure)
	v, err := strconv.ParseInt(sv, 10, 32)
	if err != nil && len(defaultVs) > 0 {
		return defaultVs[0], ""
	}
	utee.Chk(err)
	return int(v), pure
}

func form_func_Int64(r *http.Request, name string) (result int64, pure string) {
	pure = r.PostFormValue(name)
	sv := strings.TrimSpace(pure)
	v, err := strconv.ParseInt(sv, 10, 64)
	utee.Chk(err)
	return v, pure
}

func form_func_String(r *http.Request, name string) (result string, pure string) {
	pure = r.PostFormValue(name)
	return strings.TrimSpace(pure), pure
}

func form_func_StrArray(r *http.Request, name, separate string) (result []string, pure string) {
	pure = r.PostFormValue(name)
	sv := strings.TrimSpace(pure)
	return strings.Split(sv, separate), pure
}

func form_func_IntArray(r *http.Request, name, separate string) (result []int, pure string) {
	pure = r.PostFormValue(name)
	sv := strings.TrimSpace(pure)
	svs := strings.Split(sv, separate)
	ivs := make([]int, 0)
	for _, v := range svs {
		if v == "" {
			continue
		}
		if iv, err := strconv.ParseInt(v, 10, 32); err == nil {
			ivs = append(ivs, int(iv))
		}
	}
	return ivs, pure
}

func form_func_Float32(r *http.Request, name string, defaultVs ...float32) (result float32, pure string) {
	pure = r.PostFormValue(name)
	sv := strings.TrimSpace(pure)
	v, err := strconv.ParseFloat(sv, 32)
	if err != nil && len(defaultVs) > 0 {
		return defaultVs[0], pure
	}
	utee.Chk(err)
	return float32(v), pure
}
