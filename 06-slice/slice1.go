package main

import (
	"fmt"
)

func main() {
	// 固定长度的数组
	var array1 [10] int
	for i := 0; i < len(array1); i++ {
		array1[i] = i
	}

	array2 := [10] int { 1, 2, 3, 4 }
	for index, value := range array2 {
		fmt.Println("index =", index, ", value =", value)
	}
}
