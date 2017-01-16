package Figo

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strings"
)

type SQLTpl struct {
	AllCount        string
	AllQuery        string
	AllQueryPage    string
	FilterCount     string
	FilterQuery     string
	FilterQueryPage string
}

func GetTableName(v interface{}) string {
	name := TypeOf(v).String()
	parser := &Parser{
		PrepareReg: []string{},
		ProcessReg: []string{"\\*", ".+\\."},
	}
	name = parser.Exe(name)[0]
	return SnakeString(name)
}

func GenSql(v interface{}) SQLTpl {
	tableName := GetTableName(v)
	return SQLTpl{
		AllCount:        fmt.Sprint("select count(*) from ", tableName),
		AllQuery:        fmt.Sprint("select * from ", tableName),
		AllQueryPage:    fmt.Sprint("select * from ", tableName, " limit #{start},#{limit}"),
		FilterCount:     fmt.Sprint("select count(*) from ", tableName, " where 1=1 and #{filter_statement}"),
		FilterQuery:     fmt.Sprint("select * from ", tableName, " where 1=1 and #{filter_statement}"),
		FilterQueryPage: fmt.Sprint("select * from ", tableName, " where 1=1 and #{filter_statement} limit #{start},#{limit}"),
	}
}

type DAO interface {
	GetORM() orm.Ormer
	GetItemContainer() interface{}
	GetItemsContainer() interface{}
}

type Manager struct {
	Dao DAO
}

func (p *Manager) QueryOne(filter string) interface{} {
	result := p.Dao.GetItemContainer()
	tableName := GetTableName(result)
	dao := p.Dao.GetORM()
	query := fmt.Sprint("select * from ", tableName, " where 1=1 and ", filter)
	log.Println("@Query:", query)
	dao.Raw(query).QueryRow(result)
	return result
}

func (p *Manager) QueryList(filter string) interface{} {
	result := p.Dao.GetItemsContainer()
	tableName := GetTableName(p.Dao.GetItemContainer())
	dao := p.Dao.GetORM()
	query := fmt.Sprint("select * from ", tableName, " where 1=1 and ", filter)
	dao.Raw(query).QueryRow(&result)
	return result
}

func (p *Manager) CountAll() int {
	dao := p.Dao.GetORM()
	sqlTpl := GenSql(p.Dao.GetItemContainer())
	var count int
	dao.Raw(sqlTpl.AllCount).QueryRow(&count)
	return count
}

func (p *Manager) QueryAll() interface{} {
	dao := p.Dao.GetORM()
	sqlTpl := GenSql(p.Dao.GetItemContainer())
	data := p.Dao.GetItemsContainer()
	dao.Raw(sqlTpl.AllQuery).QueryRows(data)
	return data
}

func (p *Manager) QueryAllPaging(start, limit int) interface{} {
	dao := p.Dao.GetORM()
	sqlTpl := GenSql(p.Dao.GetItemContainer())
	query := strings.Replace(sqlTpl.AllQueryPage, "#{start}", fmt.Sprint(start), -1)
	query = strings.Replace(query, "#{limit}", fmt.Sprint(limit), -1)
	data := p.Dao.GetItemsContainer()
	dao.Raw(query).QueryRows(data)
	return data
}

func (p *Manager) CountFilter(filter string) int {
	dao := p.Dao.GetORM()
	sqlTpl := GenSql(p.Dao.GetItemContainer())
	var count int
	query := strings.Replace(sqlTpl.FilterCount, "#{filter_statement}", filter, -1)
	dao.Raw(query).QueryRow(&count)
	return count
}

func (p *Manager) QueryFilter(filter string) interface{} {
	dao := p.Dao.GetORM()
	sqlTpl := GenSql(p.Dao.GetItemContainer())
	data := p.Dao.GetItemsContainer()
	query := strings.Replace(sqlTpl.FilterQuery, "#{filter_statement}", filter, -1)
	dao.Raw(query).QueryRows(data)
	return data
}

func (p *Manager) QueryFilterPaging(filter string, start, limit int) interface{} {
	dao := p.Dao.GetORM()
	sqlTpl := GenSql(p.Dao.GetItemContainer())
	query := strings.Replace(sqlTpl.FilterQueryPage, "#{start}", fmt.Sprint(start), -1)
	query = strings.Replace(query, "#{limit}", fmt.Sprint(limit), -1)
	query = strings.Replace(query, "#{filter_statement}", filter, -1)
	data := p.Dao.GetItemsContainer()
	dao.Raw(query).QueryRows(data)
	return data
}
