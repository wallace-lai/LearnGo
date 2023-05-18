// 多维动态规划

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/minimum-path-sum/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	path := make([][]int, m + 1, m + 1)
	for i := 0; i < m + 1; i++ {
		path[i] = make([]int, n + 1, n + 1)
	}
	for i := 0; i < m + 1; i++ {
		path[i][0] = math.MaxInt32
	}
	for j := 0; j < n + 1; j++ {
		path[0][j] = math.MaxInt32
	}

	for i := 1; i < m + 1; i++ {
		for j := 1; j < n + 1; j++ {
			if i == 1 && j == 1 {
				path[i][j] = grid[i - 1][j - 1]
				continue
			}

			path[i][j] = getMin(path[i][j - 1], path[i - 1][j]) + grid[i - 1][j - 1]
		}
	}

	return path[m][n]
}

func main() {
	grid := [][]int {
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	fmt.Println(minPathSum(grid))
}