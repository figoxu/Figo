package Figo

import (
	"expvar"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/quexer/utee"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func NewMartini(slowMSec, concurrent int, war string) *martini.ClassicMartini {
	m := martini.Classic()
	m.Handlers(martini.Recovery())
	m.Use(MidSlowLog(slowMSec))
	m.Use(MidConcurrent(concurrent))
	m.Use(martini.Static(war, martini.StaticOptions{SkipLogging: true}))
	http.Handle("/", m)
	return m
}

type tmEntry struct {
	name string
	tick int64
}
type TimeMatrix struct {
	sync.Mutex
	m []*tmEntry
}

func (p *TimeMatrix) Rec(name string) {
	p.Lock()
	p.m = append(p.m, &tmEntry{name, utee.Tick()})
	p.Unlock()
}

func (p *TimeMatrix) Print() {
	p.Lock()
	if len(p.m) < 2 {
		return
	}

	l := []string{"time matrix"}
	for i, val := range p.m {
		if i == 0 {
			continue
		}
		s := fmt.Sprintf("%31v: %5vms", p.m[i-1].name+"~"+val.name, val.tick-p.m[i-1].tick)
		l = append(l, s)
	}
	log.Println(strings.Join(l, "\n"))
	p.Unlock()
}

func MidSlowLog(limit int) func(*http.Request, martini.Context) {
	if limit <= 0 {
		log.Fatalln("[slow log] err:  bad limit")
	}

	return func(req *http.Request, c martini.Context) {
		start := utee.Tick()
		tm := &TimeMatrix{}
		tm.Rec("start")
		c.Map(tm)
		defer func() {
			t := utee.Tick() - start
			if t >= int64(limit) {
				log.Printf("[slow] %3vms %s \n", t, req.RequestURI)
				if utee.Env("TIME_MATRIX", false, false) != "" {
					tm.Print()
				}
			}
		}()
		c.Next()
		tm.Rec("end")
	}
}

var (
	expServerConcurrent = expvar.NewInt("z_utee_server_concurrent")
	expServeCount       = expvar.NewInt("z_utee_serve_count")
	expTps              = expvar.NewInt("z_utee_serve_tps")
)

func MidConcurrent(concurrent ...int) func(http.ResponseWriter, martini.Context) {
	go func() {
		lastSecond, _ := strconv.ParseInt(expServeCount.String(), 10, 64)
		for range time.Tick(time.Second) {
			countTotal, _ := strconv.ParseInt(expServeCount.String(), 10, 64)
			expTps.Set(countTotal - lastSecond)
			lastSecond = countTotal
		}
	}()

	var ch chan string
	if len(concurrent) > 0 {
		if concurrent[0] > 0 {
			ch = make(chan string, concurrent[0])
		} else {
			log.Fatalln("bad concurrent number", concurrent[0])
		}
	}
	return func(w http.ResponseWriter, c martini.Context) {
		expServerConcurrent.Add(1)
		defer func() {
			expServerConcurrent.Add(-1)
			expServeCount.Add(1)
		}()
		if ch == nil {
			c.Next()
		} else {
			select {
			case ch <- "a":
				func() {
					defer func() {
						<-ch
					}()
					c.Next()
				}()
			default:
				log.Println("[warn] reach concurrent limit:", concurrent[0])
				http.Error(w, "server is busy", 503)
			}
		}
	}
}
