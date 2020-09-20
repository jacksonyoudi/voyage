package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4}

	fmt.Println(arr1)
	fmt.Println(arr2)

	slice := []int{1, 2, 3}
	slice2 := make([]int, 10)

	// append
	fmt.Println(slice)
	fmt.Println(slice2)

	hash := make(map[string]int, 3)
	hash["1"] = 2
	hash["3"] = 4
	hash["5"] = 6

}
