// 矩阵

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/set-matrix-zeroes/?envType=study-plan-v2&id=top-interview-150

// v1 : 空间复杂度 O(m + n)，8ms
func setZeroes(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])

	// rows和cols标记对应的行与列是否需要置零
	rows := make([]int, m)
	cols := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = 1
				cols[j] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] == 1 || cols[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int {
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(matrix)
	fmt.Println(matrix)
}