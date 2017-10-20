package Figo

import (
	"sync"
	"github.com/quexer/utee"
)

type BlockChannelItem struct {
	k        interface{}
	data     interface{}
	done     chan bool
	doneFlag bool
	tryTimes int
}
// 需要仔细思考 doneFlag是否解决 异步乱序、导致重复执行的问题

func (p *BlockChannelItem) timesIncr() {
	p.tryTimes = p.tryTimes + 1
}

type BlockExecuteQ struct {
	q        utee.MemQueue
	tc       *utee.TimerCache
	seq      *SeqMem
	tryTimes int
	execute  func(interface{}, chan bool)
	mutex    sync.Mutex
}
// 使用mutex 避免 retry 和 blockExc 乱序执行

func NewBlockExecuteQ(cap, retrySec, tryTimes int, exec func(interface{}, chan bool)) BlockExecuteQ {
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
		p.tc.Put(k, v)
		p.execute(task.data, task.done)
	}
}

func (p *BlockExecuteQ) blockExec(v interface{}) {
	task := v.(*BlockChannelItem)
	p.tc.Put(task.k, task)
	task.timesIncr()
	p.execute(task.data, task.done)
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
