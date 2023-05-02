// 矩阵

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/rotate-image/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func rotate(matrix [][]int) {
	n := len(matrix)

	for x := 0; x < n; x++ {
		for y := 0; y < x; y++ {
			tmp := matrix[x][y]
			matrix[x][y] = matrix[y][x]
			matrix[y][x] = tmp
		}
	}

	for y := 0; y < n/2; y++ {
		for x := 0; x < n; x++ {
			tmp := matrix[x][y]
			matrix[x][y] = matrix[x][n-y-1]
			matrix[x][n-y-1] = tmp
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)
	fmt.Println(matrix)
}
