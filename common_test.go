package Figo

import (
	"errors"
	"fmt"
	"github.com/quexer/utee"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCatch(t *testing.T) {
	log.Println("Before")
	catchSampleMethod()
	log.Println("After")
}

func catchSampleMethod() {
	defer Catch()
	log.Println("Hello")
	utee.Chk(errors.New("Some Error"))
}

type SubCloneObj struct {
	sync.RWMutex
	data  map[string]string
	key   string
	score int
}

type CloneObj struct {
	sync.RWMutex
	Leader   SubCloneObj
	Member   []SubCloneObj
	TeamName string
	Year     int
}

func (p *CloneObj) display() {
	fmt.Println("@Year:", p.Year, " @TeamName:", p.TeamName, "@Leader:", p.Leader)
	for _, obj := range p.Member {
		fmt.Println(obj)
	}
}

func (p *CloneObj) AppendChild(objs []SubCloneObj) {
	log.Println("===>")
	log.Println(p.Member)
	log.Println("<===")
	p.Member = append(p.Member, objs...)
}

func RandomCloneObj() SubCloneObj {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rv := func(v string) string {
		return fmt.Sprint(v, "-", r.Int())
	}
	obj := SubCloneObj{
		data:  map[string]string{rv("foo"): rv("bar"), rv("foo1"): rv("bar1"), rv("foo2"): rv("bar2")},
		key:   rv("Figo"),
		score: r.Intn(100),
	}
	return obj
}

func RandomCloneObjs(length int) []SubCloneObj {
	objs := []SubCloneObj{}
	for i := 0; i < length; i++ {
		objs = append(objs, RandomCloneObj())
	}
	return objs
}

func TestClone(t *testing.T) {
	cloneObj := CloneObj{
		Leader:   RandomCloneObj(),
		TeamName: "Bee",
		Year:     2016,
		Member:   RandomCloneObjs(1),
	}
	obj2 := Clone(cloneObj).(CloneObj)
	obj2.AppendChild(RandomCloneObjs(2))
	obj2.Year = 2017

	log.Println("===============")
	cloneObj.display()
	log.Println("===============")
	obj2.display()

}
