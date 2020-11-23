package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func loadNewConfig() Config {
	return Config{
		"北京",
		"10.77.95.27",
		rand.Int31(),
	}
}

func main() {
	var config atomic.Value
	config.Store(loadNewConfig())
	var cond = sync.NewCond(&sync.Mutex{})

	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast()
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			cond.Wait()
			c := config.Load().(Config)
			println("config", c)
			cond.L.Lock()
		}
	}()

	select {}

}
