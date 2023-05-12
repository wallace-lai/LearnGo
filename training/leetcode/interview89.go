// 图

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/number-of-islands/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func depthFirstSearch(grid *[][]byte, m, n, i, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || (*grid)[i][j] != '1' {
		return
	}

	(*grid)[i][j] = '0'

	depthFirstSearch(grid, m, n, i - 1, j)
	depthFirstSearch(grid, m, n, i + 1, j)
	depthFirstSearch(grid, m, n, i, j - 1)
	depthFirstSearch(grid, m, n, i, j + 1)
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	result := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				result++
				depthFirstSearch(&grid, m, n, i, j)
			}
		}
	}
	return result
}

