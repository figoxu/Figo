package Figo

import (
	"log"
	"errors"
)

var ErrFull = errors.New("queue is full")

type BatchMemQueue chan interface{}

func NewBatchMemQueue(cap, concurrent, batchSize int, worker func([]interface{})) BatchMemQueue {
	var q BatchMemQueue
	q = make(chan interface{}, cap)
	f := func() {
		for {
			worker(q.DeqN(batchSize))
		}
	}
	for i := 0; i < concurrent; i++ {
		go f()
	}
	return q
}

func (p BatchMemQueue) DeqN(n int) []interface{} {
	if n <= 0 {
		log.Println("[MemQueue] deqn err, n must > 0")
		return nil
	}

	var l []interface{}

	for {
		select {
		case data := <-p:
			l = append(l, data)
			if len(l) == n {
				return l
			}
		default:
			return l
		}
	}
}

func (p BatchMemQueue) Enq(data interface{}) error {
	select {
	case p <- data:
	default:
		return ErrFull
	}
	return nil
}
