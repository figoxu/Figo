package Figo

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"time"
	"strings"
	"strconv"
	"github.com/quexer/utee"
)

type ParamHelper struct {
	Float64 func(name string) float64
	Bool    func(name string) bool
	Int     func(name string, defaultVs ...int) int
	Int64   func(name string) int64
	Uint64  func(name string) uint64
	Time    func(name, format string) time.Time
	TimeLoc func(name, format string, loc *time.Location) time.Time
	String  func(name string) string
	IntArr  func(name, separate string) []int
}

func Mid_helper_param(c martini.Context, param martini.Params, w http.ResponseWriter) {
	c.Map(&ParamHelper{
		Float64: wp_func_float64(param),
		Bool:    wp_func_Bool(param),
		Int:     wp_func_Int(param),
		Int64:   wp_func_Int64(param),
		Uint64:  wp_func_Uint64(param),
		Time:    wp_func_time(param),
		TimeLoc: wp_func_time_loc(param),
		String:  wp_func_string(param),
	})
}

type FormHelper struct {
	Int     func(name string, defaultVs ...int) int
	Int64   func(name string) int64
	Float32 func(name string, defaultVs ...float32) float32
	String  func(name string) string
	StrArr  func(name, separate string) []string
	IntArr  func(name, separate string) []int
}

func Mid_helper_form(c martini.Context, r *http.Request) {
	c.Map(&FormHelper{
		Int:     form_func_Int(r),
		Int64:   form_func_Int64(r),
		Float32: form_func_Float32(r),
		String:  form_func_String(r),
		StrArr:  form_func_StrArray(r),
		IntArr:  form_func_IntArray(r),
	})
}

func wp_func_string(param martini.Params) func(name string) string {
	return func(name string) string {
		return strings.TrimSpace(param[name])
	}
}

func wp_func_float64(param martini.Params) func(name string) float64 {
	return func(name string) float64 {
		log.Println("wp_func_float64  ", name)
		v, err := strconv.ParseFloat(param[name], 64)
		utee.Chk(err)
		return v
	}
}

func wp_func_Bool(param martini.Params) func(name string) bool {
	return func(name string) bool {
		log.Println("wp_func_Bool  ", name)
		v, err := strconv.ParseBool(param[name])
		utee.Chk(err)
		return v
	}
}

func wp_func_Int(param martini.Params) func(name string, defaultVs ...int) int {
	return func(name string, defaultVs ...int) int {
		log.Println("wp_func_Int  ", name, " @val: ", param[name])
		v, err := strconv.ParseInt(param[name], 10, 32)
		if err != nil && len(defaultVs) > 0 {
			return defaultVs[0]
		}
		utee.Chk(err)
		return int(v)
	}
}

func wp_func_Int64(param martini.Params) func(name string) int64 {
	return func(name string) int64 {
		v, err := strconv.ParseInt(param[name], 10, 32)
		utee.Chk(err)
		return v
	}
}

func wp_func_Uint64(param martini.Params) func(name string) uint64 {
	return func(name string) uint64 {
		v, err := TpUint64(param[name])
		utee.Chk(err)
		return v
	}
}

func wp_func_time(param martini.Params) func(name, format string) time.Time {
	return func(name, format string) time.Time {
		log.Println("wp_func_time   ", name)
		t, err := time.Parse(format, param[name])
		utee.Chk(err)
		return t
	}
}

func wp_func_time_loc(param martini.Params) func(name, format string, loc *time.Location) time.Time {
	return func(name, format string, loc *time.Location) time.Time {
		log.Println("wp_func_time   ", name)
		t, err := time.ParseInLocation(format, param[name], local)
		utee.Chk(err)
		return t
	}
}

func wp_func_IntArray(param martini.Params) func(name, separate string) []int {
	return func(name, separate string) []int {
		sv := strings.TrimSpace(param[name])
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
		return ivs
	}
}

func form_func_Int(r *http.Request) func(name string, defaultVs ...int) int {
	return func(name string, defaultVs ...int) int {
		sv := strings.TrimSpace(r.PostFormValue(name))
		v, err := strconv.ParseInt(sv, 10, 32)
		if err != nil && len(defaultVs) > 0 {
			return defaultVs[0]
		}
		utee.Chk(err)
		return int(v)
	}
}

func form_func_Int64(r *http.Request) func(name string) int64 {
	return func(name string) int64 {
		sv := strings.TrimSpace(r.PostFormValue(name))
		v, err := strconv.ParseInt(sv, 10, 32)
		utee.Chk(err)
		return v
	}
}

func form_func_String(r *http.Request) func(name string) string {
	return func(name string) string {
		return strings.TrimSpace(r.PostFormValue(name))
	}
}

func form_func_StrArray(r *http.Request) func(name, separate string) []string {
	return func(name, separate string) []string {
		sv := strings.TrimSpace(r.PostFormValue(name))
		return strings.Split(sv, separate)
	}
}

func form_func_IntArray(r *http.Request) func(name, separate string) []int {
	return func(name, separate string) []int {
		sv := strings.TrimSpace(r.PostFormValue(name))
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
		return ivs
	}
}

func form_func_Float32(r *http.Request) func(name string, defaultVs ...float32) float32 {
	return func(name string, defaultVs ...float32) float32 {
		sv := strings.TrimSpace(r.PostFormValue(name))
		v, err := strconv.ParseFloat(sv, 32)
		if err != nil && len(defaultVs) > 0 {
			return defaultVs[0]
		}
		utee.Chk(err)
		return float32(v)
	}
}
