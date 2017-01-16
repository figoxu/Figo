package Figo

import (
	"errors"
	"fmt"
	"github.com/quexer/utee"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestCatch(t *testing.T) {
	log.Println("Before")
	catchSampleMethod()
	log.Println("After")
	log.Println("Panic In Loops Begin")
	for i := 0; i < 100; i++ {
		{
			defer Catch()
			utee.Chk(errors.New(fmt.Sprint("error", i)))
		}
	}
	log.Println("Panic In Loops Finish")
}

func catchSampleMethod() {
	defer Catch()
	log.Println("Hello")
	utee.Chk(errors.New("Some Error"))
}

type SubCloneObj struct {
	data  map[string]string
	key   string
	score int
}

type CloneObj struct {
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

	//	log.Println("1==>")
	//	 Clone(&cloneObj)
	//	log.Println("2==>")
	Clone(cloneObj)
}

func TestExist(t *testing.T) {
	v := Exist(1, 2, 3, 4, 1)
	log.Println(v)
	a := []string{"hello", "word"}
	log.Println(Exist("hello", ICast.FromStrings(a)...))
	log.Println(Exist("foo", ICast.FromStrings(a)...))
}

func TestRetryExe(t *testing.T) {
	alwaysError := func() error {
		return errors.New("test")
	}
	RetryExe(alwaysError, 3, " alwaysError() Method")
	successFunc := func() error {
		log.Println("execute")
		return nil
	}
	RetryExe(successFunc, 3, " successFunc() Method")
	tryCount := 0
	retrySuccessFunc := func() error {
		tryCount++
		if tryCount >= 2 {
			return nil
		}
		return errors.New("retry error")
	}
	RetryExe(retrySuccessFunc, 3, " retrySuccessFunc() Method")
}

func TestPrint(t *testing.T) {
	Print(THEME_Magenta, "how ")
	Print(THEME_Blue, "are ")
	Print(THEME_Green, "you ")
	Println(THEME_Magenta, "hello")
	log.Println("hello")
}
