package main

import "fmt"

/**
! 匿名函数
! 赋值给一个变量
! 全局匿名函数
*/

var (
	Func1 = func(a, b int) int {
		return a + b
	}
)

func main() {

	c := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(c)

	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(4, 5))

	fmt.Println(Func1(10, 11))

}
