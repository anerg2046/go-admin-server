package concurrency

import "sync"

type Pool struct {
	ch chan bool
	wg *sync.WaitGroup
}

func New(count int) *Pool {
	if count < 1 {
		count = 1
	}
	return &Pool{
		ch: make(chan bool, count),
		wg: new(sync.WaitGroup),
	}
}

func (c *Pool) Add() {
	c.ch <- true
	c.wg.Add(1)
}

func (c *Pool) Done() {
	<-c.ch
	c.wg.Done()
}

func (c *Pool) Wait() {
	c.wg.Wait()
}
