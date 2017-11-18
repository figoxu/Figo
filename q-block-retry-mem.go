package Figo

import (
	"github.com/quexer/utee"
	"sync"
)

type BlockChannelItem struct {
	k        interface{}
	data     interface{}
	done     chan bool
	doneFlag bool
	tryTimes int
}

func (p *BlockChannelItem) timesIncr() {
	p.tryTimes = p.tryTimes + 1
}

type BlockExecuteQ struct {
	q        utee.MemQueue
	tc       *utee.TimerCache
	seq      *SeqMem
	tryTimes int
	execute  func(interface{}) bool
	mutex    sync.Mutex
}

func NewBlockExecuteQ(cap, retrySec, tryTimes int, exec func(interface{}) bool) BlockExecuteQ {
	beq := BlockExecuteQ{
		tryTimes: tryTimes,
		execute:  exec,
		seq:      NewSeqMem(),
		mutex:    sync.Mutex{},
	}
	q := utee.NewLeakMemQueue(cap, 1, beq.blockExec)
	tc := utee.NewTimerCache(retrySec, beq.retry)
	beq.q = q
	beq.tc = tc
	return beq
}

func (p *BlockExecuteQ) retry(k, v interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	task := v.(*BlockChannelItem)
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

func (p *BlockExecuteQ) blockExec(v interface{}) {
	task := v.(*BlockChannelItem)
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

func (p *BlockExecuteQ) Enq(v interface{}) {
	k := byte(p.seq.Next())
	item := &BlockChannelItem{
		k:        k,
		data:     v,
		done:     make(chan bool, 1),
		tryTimes: 0,
		doneFlag: false,
	}
	p.q.Enq(item)
}
