package main

import "fmt"

func main() {
	n := 1000
	fmt.Printf("%T, %v, %v\n", n, n, &n)

	n1 := new(int) // 指针类型
	fmt.Printf("%T, %v, %v， %v", n1, n1, &n1, *n1)


	// make 分配内存， 主要用于引用类型的

}
