package Figo

import (
	"gopkg.in/mgo.v2"
	"log"
	"sync"
)

type MgoSe struct {
	sync.Mutex
	realSession  *mgo.Session
	referSession *mgo.Session
}

func NewMgoSe(referSession *mgo.Session) *MgoSe {
	return &MgoSe{
		referSession: referSession,
	}
}
func (p *MgoSe) Copy() *MgoSe {
	return NewMgoSe(p.referSession)
}

func (p *MgoSe) SetMode(consistency mgo.Mode, refresh bool) {
	p.Session().SetMode(consistency, refresh)
}

func (p *MgoSe) DB(name string) *mgo.Database {
	return p.Session().DB(name)
}

func (p *MgoSe) IsOpen() bool {
	if p.realSession == nil {
		return false
	}
	return true
}

func (p *MgoSe) Session() *mgo.Session {
	p.Lock()
	defer p.Unlock()
	if !p.IsOpen() {
		if p.referSession == nil {
			log.Panicln("error : referSession is not configure")
		}
		p.realSession = p.referSession.Copy()
	}
	return p.realSession
}

func (p *MgoSe) Close() {
	p.Lock()
	defer p.Unlock()
	if p.IsOpen() {
		p.realSession.Close()
		p.realSession = nil
	}
}
