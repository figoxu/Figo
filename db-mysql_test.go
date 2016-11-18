package Figo

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type Test1 struct {
	Id   int
	Foo  string
	Bar  string
	Info string `orm:"default(Hello FIgo)"`
}

type Test2 struct {
	Id     int
	Hello  string
	Salary int
	Age    int `orm:"default(32)"`
}

func TestMysqlConf(t *testing.T) {
	conf := MysqlConf{

		User:       "root",
		Pwd:        "4rfv%TGB",
		Host:       "115.159.104.88",
		Port:       "3306",
		Name:       "figo_research",
		ConnIdle:   1,
		ConnActive: 1,
	}

	conf.Conf(new(Test1), new(Test2))
	orm.RunSyncdb("default", false, true)
}
