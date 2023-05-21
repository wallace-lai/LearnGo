// 数学

package main

import "fmt"

// https://leetcode.cn/problems/factorial-trailing-zeroes/?envType=study-plan-v2&envId=top-interview-150

// v1 : 4ms
// func trailingZeroes(n int) int {
// 	result := 0
// 	for i := 5; i < n + 1; i = i + 5 {
// 		for x := i; x % 5 == 0; x /= 5 {
// 			result++
// 		}
// 	}
// 	return result
// }

// v2 : 0ms
func trailingZeroes(n int) int {
	result := 0
	for n != 0 {
		n = n / 5
		result += n
	}
	return result
}

func main() {
	fmt.Println(trailingZeroes(10))
}