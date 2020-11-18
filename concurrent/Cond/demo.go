package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			cond.L.Lock()
			ready++
			cond.L.Unlock()


			log.Println("运动员#%d 已准备就绪")
			cond.Broadcast()
		}(i)
	}

	cond.L.Lock()
	for ready != 10 {
		cond.Wait()
		log.Println("裁判员被唤醒一次")
	}
	cond.L.Unlock()

	log.Println("所有运动员都准备就绪。比赛开始 3，2，1")

}
