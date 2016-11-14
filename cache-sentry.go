package Figo

import "log"

type Sentry struct {
	caches []Cache
	kpi_es []CacheKPI
	notify func(key interface{}) error
}

func NewSentry(notify func(key interface{}) error, caches ...Cache) *Sentry {
	kpi_es := []CacheKPI{}
	for index, _ := range caches {
		log.Println("init cache level:", index)
		kpi_es = append(kpi_es, CacheKPI{})
	}
	return &Sentry{
		caches: caches,
		notify: notify,
		kpi_es: kpi_es,
	}
}

func (p *Sentry) Put(key, value interface{}) {
	for _, cache := range p.caches {
		cache.Put(key, value)
	}
	if p.notify != nil {
		go p.notify(key)
	}
}

func (p *Sentry) Get(key interface{}) interface{} {
	upgradeCache := func(key, value interface{}, caches ...Cache) {
		for _, cache := range caches {
			cache.Put(key, value)
		}
	}
	missCaches := []Cache{}
	for index, cache := range p.caches {
		kpi := p.kpi_es[index]
		if v := cache.Get(key); v != nil {
			kpi.Save(true)
			go upgradeCache(key, v, missCaches...)
			return v
		} else {
			kpi.Save(false)
			missCaches = append(missCaches, cache)
		}
	}
	return nil
}
