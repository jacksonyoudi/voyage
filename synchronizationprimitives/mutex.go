package main

import "sync"

func main() {
	var mutex sync.Mutex

	mutex.Lock()
}
