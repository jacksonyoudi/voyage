package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter CounterD
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				counter.Incr()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Counter())
}

type CounterD struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

// 自增 需要加锁
func (c *CounterD) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *CounterD) Counter() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
