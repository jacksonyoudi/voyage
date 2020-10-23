package main

import "sync"

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}

		}()
	}

}

// 使用嵌入字段的方式
type Counter struct {
	sync.Mutex
	Count uint64
}
