// 数学

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/sqrtx/?envType=study-plan-v2&envId=top-interview-150

// v1 : 0ms
// func mySqrt(x int) int {
// 	if x == 1 {
// 		return 1
// 	}

// 	mid := 0
// 	left := 0
// 	right := x

// 	for right - left > 1 {
// 		mid = left + (right - left) / 2
// 		if mid > x / mid {
// 			right = mid
// 		} else {
// 			left = mid
// 		}
// 	}
// 	return left
// }

// v2 : 4ms
// func mySqrt(x int) int {
// 	result := -1

// 	mid := 0
// 	left := 0
// 	right := x
// 	for left <= right {
// 		mid = left + (right - left) / 2
// 		if mid * mid <= x {
// 			result = mid
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	return result
// }

// v3 : 0ms
func mySqrt(x int) int {
	if x == 0 {
		return x
	}

	result := int(math.Exp(0.5 * math.Log(float64(x))))
	if (result + 1) * (result + 1) <= x {
		return result + 1
	}
	return result
}


func main() {
	x := 1
	fmt.Println(mySqrt(x))
}

