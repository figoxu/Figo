package Figo

import (
	"github.com/jinzhu/gorm"
	"log"
)

type GormLog struct {
}

// @index: 0    @value: sql
// @index: 1    @value: /Users/xujianhui/GoglandProjects/sdz-stock-service/dao/db-CampSite.go:117
// @index: 2    @value: 106.335985ms
// @index: 3    @value:  select * from camp_site order by id
// @index: 4    @value: []   => result type
// @index: 5    @value: 102
func (p *GormLog) Print(values ...interface{}) {
	log.Println(gorm.LogFormatter(values...)...)
}
