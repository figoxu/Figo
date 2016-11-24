package Figo

type Cache interface {
	Put(key, val interface{})
	Get(key interface{}) interface{}
}

type CacheObj struct {
	put func(key, val interface{})
	get func(key interface{}) interface{}
}

func (p *CacheObj) Put(key, val interface{}) {
	p.put(key, val)
}

func (p *CacheObj) Get(key interface{}) interface{} {
	return p.get(key)
}

func NewCacheObj(put func(key, val interface{}), get func(key interface{}) interface{}) *CacheObj {
	return &CacheObj{
		put: put,
		get: get,
	}
}
