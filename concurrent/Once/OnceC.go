package main

import (
	"sync"
	"sync/atomic"
)

type OnceC struct {
	done uint32
	mu   sync.Mutex
}

func (c *OnceC) Do(f func()) {
	if atomic.LoadUint32(&c.done) == 0 {
		c.doSlow(f)
	}
}

func (c *OnceC) doSlow(f func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.done == 0 {
		defer atomic.StoreUint32(&c.done, 1)
		f()
	}

}
