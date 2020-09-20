package main

import "fmt"

type MyStruct struct {
	i int
}

func my_Function(a MyStruct, b *MyStruct) {
	a.i = 31
	b.i = 41
	fmt.Printf("in my_function - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}
func main() {
	a := MyStruct{i: 30}
	b := &MyStruct{i: 40}
	fmt.Printf("before calling - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
	my_Function(a, b)
	fmt.Printf("after calling  - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}
