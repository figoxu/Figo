package Figo

import (
	"github.com/quexer/utee"
)

type BothWayMap struct {
	ab  map[interface{}]interface{}
	ba  map[interface{}]interface{}
	ttl *utee.TimerCache
}

func NewBothWayMap(ttlTime int) *BothWayMap {
	bothWayMap := &BothWayMap{
		ab: make(map[interface{}]interface{}),
		ba: make(map[interface{}]interface{}),
	}
	if ttlTime > 0 {
		bothWayMap.ttl = utee.NewTimerCache(ttlTime, bothWayMap.ttlCb)
	}
	return bothWayMap
}

func (p *BothWayMap) ttlCb(key, value interface{}) {
	p.DeleteKey(key)
}

func (p *BothWayMap) Put(key, value interface{}) *BothWayMap {
	p.ab[key] = value
	p.ba[value] = key
	if p.ttl != nil {
		p.ttl.Put(key, value)
	}
	return p
}

func (p *BothWayMap) GetByKey(key interface{}) (value interface{}, exists bool) {
	value, exists = p.ab[key]
	return
}

func (p *BothWayMap) GetByValue(value interface{}) (key interface{}, exists bool) {
	key, exists = p.ba[value]
	return
}

func (p *BothWayMap) Len() int {
	return len(p.ab)
}

func (p *BothWayMap) DeleteKey(key interface{}) *BothWayMap {
	value, exists := p.ab[key]
	if exists {
		delete(p.ab, key)
		delete(p.ba, value)
		if p.ttl != nil {
			p.ttl.Remove(key)
		}
	}
	return p
}

func (p *BothWayMap) DeleteValue(value interface{}) *BothWayMap {
	key, exists := p.ba[value]
	if exists {
		delete(p.ab, key)
		delete(p.ba, value)
		if p.ttl != nil {
			p.ttl.Remove(key)
		}
	}
	return p
}
