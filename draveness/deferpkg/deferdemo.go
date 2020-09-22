package main

import "fmt"

/**
函数执行完结束，清理收尾工作
*/

func sum(n1, n2 int) int {
	defer fmt.Println("one n1=", n1)
	// !值拷贝
	defer fmt.Println("two n2=", n2)


	n1++
	n2++
	res := n1 + n2
	fmt.Println("res=", res)

	return res

}

func main() {
	res := sum(10, 20)
	fmt.Println("sum res=", res)

}
