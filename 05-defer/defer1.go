package main

import (
	"fmt"
)

func main() {
	// 在语句前加上defer关键字
	defer fmt.Println("main end 1 !")
	defer fmt.Println("main end 2 !")

	fmt.Println("hello world !")
	fmt.Println("hello go !")
}
