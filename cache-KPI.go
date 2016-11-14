package Figo

import "log"

type CacheKPI struct {
	total int64
	miss  int64
}

func (p *CacheKPI) Save(hit bool) {
	p.total++
	if !hit {
		p.miss++
	}
}

func (p *CacheKPI) Print() {
	missRate := (p.miss * 100) / p.total
	hitRate := 100 - missRate
	log.Println("@total:", p.total, " @Hits:", (p.total - p.miss), " @HitsRate:", hitRate, "%  @Miss:", p.miss, " @MissRate:", missRate, "%")
}
