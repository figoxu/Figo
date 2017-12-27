package Figo

import (
	"github.com/quexer/utee"
)

type BothWayMap struct {
	kv  map[interface{}]interface{}
	vk  map[interface{}]interface{}
	ttl *utee.TimerCache
}

func NewBothWayMap(ttlTime int) *BothWayMap {
	bothWayMap := &BothWayMap{
		kv: make(map[interface{}]interface{}),
		vk: make(map[interface{}]interface{}),
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
	p.kv[key] = value
	p.vk[value] = key
	if p.ttl != nil {
		p.ttl.Put(key, value)
	}
	return p
}

func (p *BothWayMap) GetByKey(key interface{}) (value interface{}, exists bool) {
	value, exists = p.kv[key]
	return
}

func (p *BothWayMap) GetByValue(value interface{}) (key interface{}, exists bool) {
	key, exists = p.vk[value]
	return
}

func (p *BothWayMap) Len() int {
	return len(p.kv)
}

func (p *BothWayMap) DeleteKey(key interface{}) *BothWayMap {
	value, exists := p.kv[key]
	if exists {
		delete(p.kv, key)
		delete(p.vk, value)
		if p.ttl != nil {
			p.ttl.Remove(key)
		}
	}
	return p
}

func (p *BothWayMap) DeleteValue(value interface{}) *BothWayMap {
	key, exists := p.vk[value]
	if exists {
		delete(p.kv, key)
		delete(p.vk, value)
		if p.ttl != nil {
			p.ttl.Remove(key)
		}
	}
	return p
}
