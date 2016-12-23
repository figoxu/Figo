package Figo

import (
	"github.com/go-martini/martini"
	"github.com/quexer/utee"
	"net/http"
)

func NewMartini(slowMSec, concurrent int, war string) *martini.ClassicMartini {
	m := martini.Classic()
	m.Handlers(martini.Recovery())
	m.Use(utee.MidSlowLog(slowMSec))
	m.Use(utee.MidConcurrent(concurrent))
	m.Use(martini.Static(war, martini.StaticOptions{SkipLogging: true}))
	http.Handle("/", m)
	return m
}
