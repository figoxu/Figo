package Figo

import (
	"github.com/quexer/utee"
	"sync"
	"log"
	"time"
)

type MultiAsyncBlockExecuteQ struct {
	mq       map[string]utee.MemQueue
	mc       map[string]chan bool
	tc       *utee.TimerCache
	seq      *SeqMem
	trySecs  int
	tryTimes int
	execute  func(data interface{})
	callback func(data interface{}, workFlag bool)
	qlock    sync.Mutex
	clock    sync.Mutex
	perCap   int
}

func NewMultiAsyncBlockExecuteQ(perCap, retrySec, tryTimes int, exec func(data interface{}), callback func(data interface{}, workFlag bool)) MultiAsyncBlockExecuteQ {
	beq := MultiAsyncBlockExecuteQ{
		tryTimes: tryTimes,
		trySecs:  tryTimes,
		execute:  exec,
		callback: callback,
		seq:      NewSeqMem(),
		qlock:    sync.Mutex{},
		clock:    sync.Mutex{},
		perCap:   perCap,
	}
	beq.mq = make(map[string]utee.MemQueue)
	beq.mc = make(map[string]chan bool)
	beq.tc = utee.NewTimerCache(retrySec, beq.retry)
	return beq
}

func (p *MultiAsyncBlockExecuteQ) retry(k, v interface{}) {
	task := v.(*MultiBlockChannelItem)
	task.timesIncr()
	if task.tryTimes > p.tryTimes {
		p.tc.Remove(task.k)
		p.getC(task.prefix) <- false
		return
	}
	if !task.doneFlag {
		p.execute(task.data)
		p.tc.Put(k, v)
	}
}

func (p *MultiAsyncBlockExecuteQ) blockExec(v interface{}) {
	task := v.(*MultiBlockChannelItem)
	ch := p.getC(task.prefix)
	clearDirtyHookData := func() {
		for len(ch) > 0 {
			log.Println("clearDirtyHook >>>  @prefix:",task.prefix," @len:",len(ch))
			<-ch
		}
	}
	clearDirtyHookData()
	p.tc.Put(task.k, task)
	task.timesIncr()
	p.execute(task.data)
	workFlag:=<-ch
	task.doneFlag = true
	p.tc.Remove(task.k)
	mergeMuliResult := func(){
		for len(ch) > 0 {
			if !workFlag {
				workFlag=<-ch
			}else{
				<-ch
			}
			log.Println("merge >>>  @prefix:",task.prefix," @len:",len(ch),"@workFlag:",workFlag)
		}
	}
	time.Sleep(time.Nanosecond*time.Duration(17))
	mergeMuliResult()
	if p.callback != nil {
		p.callback(task.data,workFlag)
	}
}

func (p *MultiAsyncBlockExecuteQ) getQ(prefix string) utee.MemQueue {
	p.qlock.Lock()
	defer p.qlock.Unlock()
	if q := p.mq[prefix]; q == nil {
		p.mq[prefix] = utee.NewLeakMemQueue(p.perCap, 1, p.blockExec)
	}
	return p.mq[prefix]
}

func (p *MultiAsyncBlockExecuteQ) getC(prefix string) chan bool {
	if c := p.mc[prefix]; c != nil {
		return c
	}
	p.clock.Lock()
	defer p.clock.Unlock()
	if c := p.mc[prefix]; c != nil {
		return c
	} else {
		const WAY4Hook = 2
		p.mc[prefix] = make(chan bool, WAY4Hook)
	}
	return p.mc[prefix]
}

func (p *MultiAsyncBlockExecuteQ) Enq(prefix string, v interface{}) {
	k := p.seq.Next()
	item := &MultiBlockChannelItem{
		k:        k,
		data:     v,
		done:     make(chan bool, 1),
		tryTimes: 0,
		doneFlag: false,
		prefix:   prefix,
	}
	p.getQ(prefix).Enq(item)
}

func (p *MultiAsyncBlockExecuteQ) Deq(prefix string, filter func(item *MultiBlockChannelItem)bool) chan *MultiBlockChannelItem {
	mq:=p.getQ(prefix)
	items:=mq.DeqN(mq.Len())
	c:=make(chan *MultiBlockChannelItem,mq.Len()+1)
	for _,item:=range items {
		v:=item.(*MultiBlockChannelItem)
		if filter(v) {
			c<-v
		}
	}
	return c
}

func (p *MultiAsyncBlockExecuteQ) Hook(prefix string, status bool) {
	if status {
		p.getC(prefix) <- status
	}
}
