package Figo

import (
	"github.com/quexer/utee"
	"sync"
)

type BothWayMap struct {
	kv     map[interface{}]interface{}
	vk     map[interface{}]interface{}
	ttl    *utee.TimerCache
	rwlock sync.RWMutex
}

func NewBothWayMap(ttlTime int) *BothWayMap {
	bothWayMap := &BothWayMap{
		kv:     make(map[interface{}]interface{}),
		vk:     make(map[interface{}]interface{}),
		rwlock: sync.RWMutex{},
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
	p.rwlock.Lock()
	defer p.rwlock.Unlock()
	p.kv[key] = value
	p.vk[value] = key
	if p.ttl != nil {
		p.ttl.Put(key, value)
	}
	return p
}

func (p *BothWayMap) GetByKey(key interface{}) (value interface{}, exists bool) {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()
	value, exists = p.kv[key]
	return
}

func (p *BothWayMap) GetByValue(value interface{}) (key interface{}, exists bool) {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()
	key, exists = p.vk[value]
	return
}

func (p *BothWayMap) Len() int {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()
	return len(p.kv)
}

func (p *BothWayMap) DeleteKey(key interface{}) *BothWayMap {
	p.rwlock.Lock()
	defer p.rwlock.Unlock()
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
	p.rwlock.Lock()
	defer p.rwlock.Unlock()
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
