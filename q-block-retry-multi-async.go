package Figo

import (
	"github.com/quexer/utee"
	"sync"
	"log"
)

type MultiAsyncBlockExecuteQ struct {
	mq       map[string]utee.MemQueue
	mc       map[string]chan bool
	tc       *utee.TimerCache
	seq      *SeqMem
	trySecs  int
	tryTimes int
	execute  func(interface{})
	qlock    sync.Mutex
	clock    sync.Mutex
	perCap   int
}

func NewMultiAsyncBlockExecuteQ(perCap, retrySec, tryTimes int, exec func(v interface{})) MultiAsyncBlockExecuteQ {
	beq := MultiAsyncBlockExecuteQ{
		tryTimes: tryTimes,
		trySecs:  tryTimes,
		execute:  exec,
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
	c := p.getC(task.prefix)
	task.timesIncr()
	if task.tryTimes > p.tryTimes {
		c <- false
		p.tc.Remove(task.k)
		return
	}
	if !task.doneFlag {
		p.execute(task.data)
		p.tc.Put(k, v)
	}
}

func (p *MultiAsyncBlockExecuteQ) blockExec(v interface{}) {
	task := v.(*MultiBlockChannelItem)
	ch:=p.getC(task.prefix)
	clearDirtyHookData :=func (){
		for len(ch)>0 {
			log.Println("Clear Dirty Hook Data @prefix:",task.prefix)
			<-ch
		}
	}
	clearDirtyHookData()
	p.tc.Put(task.k, task)
	task.timesIncr()
	p.execute(task.data)

	<-ch
	task.doneFlag = true
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
	}else{
		log.Println("@new Channel @prefix:", prefix)
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

func (p *MultiAsyncBlockExecuteQ) Hook(prefix string, status bool) {
	if status {
		p.getC(prefix) <- status
	}
}
