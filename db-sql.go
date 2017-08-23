package Figo

import "fmt"

type SqlBuffer struct {
	query  string
	params []interface{}
}

func NewSqlBuffer() *SqlBuffer {
	return &SqlBuffer{
		query:  "",
		params: make([]interface{}, 0),
	}
}

func (p *SqlBuffer) Append(appendQuery string, appendParams ...interface{}) {
	p.query = fmt.Sprint(p.query, " ", appendQuery, " ")
	p.params = append(p.params, appendParams...)
}

func (p *SqlBuffer) SQL() string {
	return p.query
}

func (p *SqlBuffer) Params() []interface{} {
	return p.params
}
