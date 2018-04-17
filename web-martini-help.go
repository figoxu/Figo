package Figo

import (
	"github.com/go-martini/martini"
	"net/http"
	"time"
	"strings"
	"strconv"
	"github.com/quexer/utee"
)

type ParamHelper struct {
	context map[string]interface{}
	param   martini.Params
}

func (p *ParamHelper) Float64(name string) float64 {
	r, pure := wp_func_float64(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Bool(name string) bool {
	r, pure := wp_func_Bool(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Int(name string, defaultVs ...int) int {
	r, pure := wp_func_Int(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Int64(name string) int64 {
	r, pure := wp_func_Int64(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Uint64(name string) uint64 {
	r, pure := wp_func_Uint64(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Time(name, format string) time.Time {
	r, pure := wp_func_time(p.param, name, format)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) TimeLoc(name, format string, loc *time.Location) time.Time {
	r, pure := wp_func_time_loc(p.param, name, format, loc)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) String(name string) string {
	r, pure := wp_func_string(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) IntArr(name, separate string) []int {
	r, pure := wp_func_IntArray(p.param, name, separate)
	p.context[name] = pure
	return r
}

func Mid_helper_param(c martini.Context, param martini.Params, w http.ResponseWriter) {
	c.Map(&ParamHelper{
		param:   param,
		context: make(map[string]interface{}),
	})
}

type FormHelper struct {
	context map[string]interface{}
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

func Mid_helper_form(c martini.Context, r *http.Request) {
	c.Map(&FormHelper{
		r:       r,
		context: make(map[string]interface{}),
	})
}

func wp_func_string(param martini.Params, name string) (result string, pure string) {
	pure = param[name]
	return strings.TrimSpace(pure), pure
}

func wp_func_float64(param martini.Params, name string) (result float64, pure string) {
	pure = param[name]
	v, err := strconv.ParseFloat(pure, 64)
	utee.Chk(err)
	return v, pure
}

func wp_func_Bool(param martini.Params, name string) (result bool, pure string) {
	pure = param[name]
	v, err := strconv.ParseBool(pure)
	utee.Chk(err)
	return v, pure
}

func wp_func_Int(param martini.Params, name string, defaultVs ...int) (result int, pure string) {
	pure = param[name]
	v, err := strconv.ParseInt(pure, 10, 32)
	if err != nil && len(defaultVs) > 0 {
		return defaultVs[0], ""
	}
	utee.Chk(err)
	return int(v), pure
}

func wp_func_Int64(param martini.Params, name string) (result int64, pure string) {
	pure = param[name]
	v, err := strconv.ParseInt(pure, 10, 32)
	utee.Chk(err)
	return v, pure
}

func wp_func_Uint64(param martini.Params, name string) (result uint64, pure string) {
	pure = param[name]
	v, err := TpUint64(pure)
	utee.Chk(err)
	return v, pure
}

func wp_func_time(param martini.Params, name, format string) (result time.Time, pure string) {
	pure = param[name]
	t, err := time.Parse(format, pure)
	utee.Chk(err)
	return t, pure
}

func wp_func_time_loc(param martini.Params, name, format string, loc *time.Location) (result time.Time, pure string) {
	pure = param[name]
	t, err := time.ParseInLocation(format, pure, local)
	utee.Chk(err)
	return t, pure
}

func wp_func_IntArray(param martini.Params, name, separate string) (result []int, pure string) {
	pure = param[name]
	sv := strings.TrimSpace(pure)
	svs := strings.Split(sv, separate)
	ivs := make([]int, 0)
	for _, v := range svs {
		if v == "" {
			continue
		}
		if iv, err := strconv.ParseInt(v, 10, 32); err != nil {
			ivs = append(ivs, int(iv))
		}
	}
	return ivs, pure

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
	v, err := strconv.ParseInt(sv, 10, 32)
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
