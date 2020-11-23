package main

import "sync/atomic"

func main() {
	var a int32 = 10

	c := atomic.AddInt32(&a, 10)
	println("c:", c)
	println("a:", a)

	//
	//AddUint32(&x, ^uint32(0))

	//
	//	AddUint32(&x, ^uint32(0))

}
