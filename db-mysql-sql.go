package Figo

import (
	"fmt"
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
