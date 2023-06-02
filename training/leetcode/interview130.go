package main

import (
	"fmt"
)

// https://leetcode.cn/problems/bitwise-and-of-numbers-range/?envType=study-plan-v2&envId=top-interview-150

// v1 : 8ms
// func rangeBitwiseAnd(left int, right int) int {
// 	count := 0
// 	for left != right {
// 		left = left >> 1
// 		right = right >> 1
// 		count++
// 	}

// 	for count > 0 {
// 		left = left << 1
// 		count--
// 	}

// 	return left
// }

// v2 : 16ms
func rangeBitwiseAnd(left int, right int) int {
	for right > left {
		right = right & (right - 1)
	}

	return right
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 0))
	fmt.Println(rangeBitwiseAnd(1, 2147483647))
}
