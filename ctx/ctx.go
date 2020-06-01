package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
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
	cancel()
	time.Sleep(3 * time.Second)

}
