package main

import "fmt"

/*
闭包说明：
返回是一个匿名函数 与函数外的 环境
闭包看成一个类， 方法 与 属性
 */

func AddUpper() func(int) int {
	var n int = 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
}
