package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			// 这里可以将错误代码抛出去
		}
	}()

	n1 := 1000
	n2 := 0
	res := n1 / n2
	fmt.Println("res:", res)
}

func main() {
	test()


	fmt.Println("main() ....")
}
