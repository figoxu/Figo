package Figo

type IdService struct {
	cache Cache
	seq   Seq
}

const OFFSET_NOT_FOUND = -1

func NewIdService(cache Cache, seq Seq) *IdService {
	return &IdService{
		cache: cache,
		seq:   seq,
	}
}

func (p *IdService) GetOffSet(key string) int64 {
	if v := p.cache.Get(key); v != nil && v != OFFSET_NOT_FOUND {
		return v.(int64)
	} else {
		offset := p.seq.Next()
		p.cache.Put(key, offset)
		return offset
	}
}
