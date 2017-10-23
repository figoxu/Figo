package Figo

import (
	"github.com/quexer/utee"
	"sync"
)

type MultiBlockChannelItem struct {
	k        interface{}
	data     interface{}
	done     chan bool
	doneFlag bool
	tryTimes int
}

func (p *MultiBlockChannelItem) timesIncr() {
	p.tryTimes = p.tryTimes + 1
}

type MultiBlockExecuteQ struct {
	mq       map[string]utee.MemQueue
	tc       *utee.TimerCache
	seq      *SeqMem
	tryTimes int
	execute  func(interface{}) bool
	mutex    sync.Mutex
	qlock sync.Mutex
	perCap int
}

func NewMultiBlockExecuteQ(perCap ,retrySec, tryTimes int, exec func(interface{}) bool) MultiBlockExecuteQ {
	beq := MultiBlockExecuteQ{
		tryTimes: tryTimes,
		execute:  exec,
		seq:      NewSeqMem(),
		mutex:    sync.Mutex{},
		qlock:sync.Mutex{},
		perCap:perCap,
	}
	tc := utee.NewTimerCache(retrySec, beq.retry)
	beq.mq = make(map[string]utee.MemQueue)
	beq.tc = tc
	return beq
}

func (p *MultiBlockExecuteQ) retry(k, v interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	task := v.(*MultiBlockChannelItem)
	task.timesIncr()
	if task.tryTimes > p.tryTimes {
		task.done <- false
	} else if !task.doneFlag {
		if p.execute(task.data) {
			task.done <- true
		} else {
			p.tc.Put(k, v)
		}
	}
}

func (p *MultiBlockExecuteQ) blockExec(v interface{}) {
	task := v.(*MultiBlockChannelItem)
	p.tc.Put(task.k, task)
	task.timesIncr()
	if p.execute(task.data) {
		task.done <- true
	}
	<-task.done
	p.mutex.Lock()
	defer p.mutex.Unlock()
	task.doneFlag = true
	p.tc.Remove(task.k)
}

func (p *MultiBlockExecuteQ) getQ(prefix string) utee.MemQueue {
	p.qlock.Lock()
	defer p.qlock.Unlock()
	if q:=p.mq[prefix];q==nil {
		p.mq[prefix]=utee.NewLeakMemQueue(p.perCap, 1, p.blockExec)
	}
	return p.mq[prefix]
}

func (p *MultiBlockExecuteQ) Enq(prefix string ,v interface{}) {
	k := p.seq.Next()
	item := &MultiBlockChannelItem{
		k:        k,
		data:     v,
		done:     make(chan bool, 1),
		tryTimes: 0,
		doneFlag: false,
	}
	p.getQ(prefix).Enq(item)
}
