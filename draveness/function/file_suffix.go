package main

import (
	"fmt"
	"strings"
)

func addSuffix(s string) func(filePath string) string {
	return func(filePath string) string {
		if strings.HasSuffix(filePath, s) {
			return filePath
		} else {
			return filePath + s
		}
	}
}

func main() {
	a := addSuffix(".jpg")
	fmt.Println(a("a.jpg"))
	fmt.Println(a("b"))
}
