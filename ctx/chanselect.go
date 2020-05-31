package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("got the stop signal")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)

			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("5 s")
	stop <- true
	time.Sleep(3 * time.Second)
}
