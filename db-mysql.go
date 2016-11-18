package Figo

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConf struct {
	User       string
	Pwd        string
	Host       string
	Port       string
	Name       string
	ConnIdle   int
	ConnActive int
}

func (p *MysqlConf) Url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", p.User, p.Pwd, p.Host, p.Port, p.Name)
}

func (p *MysqlConf) Conf(models ...interface{}) {
	orm.RegisterModel(models...)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", p.Url(), p.ConnIdle, p.ConnActive)
}
