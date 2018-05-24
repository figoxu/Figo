package Figo

import (
	"os"
	"github.com/quexer/utee"
	"github.com/alicebob/qr"
)

type FileQueue struct {
	qr *qr.Qr
}

func NewFileQueue(bufferCap, concurrent int, diskLoc, queueName string, worker func(interface{})) *FileQueue {
	err := os.MkdirAll(diskLoc, 0777)
	utee.Chk(err)
	q, err := qr.New(
		diskLoc,
		queueName,
		qr.OptionBuffer(bufferCap),
	)
	utee.Chk(err)
	fq := &FileQueue{
		qr: q,
	}
	f := func() {
		for {
			worker(fq.Deq())
		}
	}
	for i := 0; i < concurrent; i++ {
		go f()
	}
	return fq
}

func (p FileQueue) Enq(data interface{}) {
	p.qr.Enqueue(data)
}

func (p FileQueue) Deq() interface{} {
	return <-p.qr.Dequeue()
}