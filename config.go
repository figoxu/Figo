package Figo

import (
	"github.com/Unknwon/goconfig"
	"sync"
)

type ConfKV struct {
	lock sync.Mutex
	db   FilePath
	conf *goconfig.ConfigFile
}

func NewConfKV(filepath string) (*ConfKV, error) {
	fp := NewFilePath(filepath)
	fullPath, err := fp.FullPath()
	if err != nil {
		return nil, err
	}
	cf, err := goconfig.LoadConfigFile(fullPath)
	if err != nil {
		return nil, err
	}
	cf.BlockMode = true
	return &ConfKV{
		db:   fp,
		conf: cf,
	}, nil
}

func (p *ConfKV) Read(id, key string) string {
	return p.conf.GetKeyComments(id, key)
}
func (p *ConfKV) ReadRecord(id string) (map[string]string, error) {
	return p.conf.GetSection(id)
}
func (p *ConfKV) Write(id, field, value string) {
	p.conf.SetValue(id, field, value)
}

func (p *ConfKV) WriteRecord(id string, record map[string]string) {
	for k, v := range record {
		p.Write(id, k, v)
	}
}

func (p *ConfKV) Flush() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	f, err := p.db.Writer()
	defer f.Flush()
	if err != nil {
		return nil
	}
	return goconfig.SaveConfigData(p.conf, f)
}
