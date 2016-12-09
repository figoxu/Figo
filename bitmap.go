package Figo

type IdService struct {
	cache Cache
	seq   Seq
}

func NewIdService(cache Cache, seq Seq) *IdService {
	return &IdService{
		cache: cache,
		seq:   seq,
	}
}

func (p *IdService) GetOffSet(key string) int64 {
	if v := p.cache.Get(key); v != nil {
		return v.(int64)
	} else {
		offset := p.seq.Next()
		p.cache.Put(key, offset)
		return offset
	}
}
