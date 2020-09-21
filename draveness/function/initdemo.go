package main

import "fmt"

/*
! 1. 如果一个文件中同时包含全局变量定义， init 函数 ，main() 函数,顺序
*/

var (
	age = test()
)

func test() int {
	fmt.Print("test")
	return 30
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Print("hello")
}
