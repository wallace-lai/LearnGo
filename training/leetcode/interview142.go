// 多维动态规划

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/triangle/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// func minimumTotal(triangle [][]int) int {
// 	n := len(triangle)
// 	if n < 2 {
// 		return triangle[0][0]
// 	}

// 	prev := make([]int, 0)
// 	curr := make([]int, 0)
// 	result := math.MaxInt32

// 	prev = append(prev, triangle[0][0])
// 	for i := 1; i < n; i++ {
// 		for j := 0; j <= i; j++ {
// 			if j == 0 {
// 				curr = append(curr, triangle[i][j] + prev[j])
// 			} else if j == i {
// 				curr = append(curr, triangle[i][j] + prev[j - 1])
// 			} else {
// 				curr = append(curr, triangle[i][j] + getMin(prev[j - 1], prev[j]))
// 			}
// 		}

// 		prev = curr
// 		curr = make([]int, 0)
// 	}

// 	for i := 0; i < len(prev); i++ {
// 		if prev[i] < result {
// 			result = prev[i]
// 		}
// 	}

// 	return result
// }

// v2 : 4ms
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}

	f := make([]int, n, n)
	f[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := i; j >= 0; j-- {
			if j == i {
				f[j] = f[j - 1] + triangle[i][j]
			} else if j == 0 {
				f[j] = f[j] + triangle[i][j]
			} else {
				f[j] = getMin(f[j - 1], f[j]) + triangle[i][j]
			}
		}
	}

	result := math.MaxInt32
	for i := 0; i < n; i++ {
		if f[i] < result {
			result = f[i]
		}
	}

	return result
}

func main() {
	triangle := [][]int {
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	fmt.Println(minimumTotal(triangle))
}