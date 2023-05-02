// 矩阵

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	// result := make([]int, m * n, m * n)
	result := make([]int, m * n)

	x, y, cnt := 0, 0, 0
	for cnt < m * n {
		// go right
		for y < n && matrix[x][y] != math.MinInt32 {
			result[cnt] = matrix[x][y]
			matrix[x][y] = math.MinInt32
			cnt++
			y++
		}
		x++
		y--

		// go down
		for x < m && matrix[x][y] != math.MinInt32 {
			result[cnt] = matrix[x][y]
			matrix[x][y] = math.MinInt32
			cnt++
			x++
		}
		x--
		y--

		// go left
		for y >= 0 && matrix[x][y] != math.MinInt32 {
			result[cnt] = matrix[x][y]
			matrix[x][y] = math.MinInt32
			cnt++
			y--
		}
		x--
		y++

		// go up
		for x >= 0 && matrix[x][y] != math.MinInt32 {
			result[cnt] = matrix[x][y]
			matrix[x][y] = math.MinInt32
			cnt++
			x--
		}
		x++
		y++
	}
	
	return result
}

func main() {
	// matrix := [][]int {
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	matrix := [][]int {
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(spiralOrder(matrix))
}