package main

import (
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {

	go func() {
		println("hello")

		id := GoID()
		println(id)

	}()
	id := GoID()
	println(id)
	time.Sleep(10 * time.Second)

}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)

	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(err)
	}
	return id
}
