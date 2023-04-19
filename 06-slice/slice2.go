package main

import (
	"fmt"
)

func printArray(array [] int) {
	// _表示匿名变量
	for _, value := range array {
		fmt.Println("value =", value)
	}
}

func main() {
	myArray := [] int { 1, 2, 3, 4 }
	fmt.Printf("myArray type is %T\n", myArray)
	printArray(myArray)
}
