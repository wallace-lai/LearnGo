// 多维动态规划

package main

// https://leetcode.cn/problems/unique-paths-ii/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func uniquePathsWithObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	path := make([][]int, m + 1, m + 1)
	for i := 0; i < m + 1; i++ {
		path[i] = make([]int, n + 1, n + 1)
	}

	for i := 1; i < m + 1; i++ {
		for j := 1; j < n + 1; j++ {
			if i == 1 && j == 1 {
				if grid[0][0] == 1 {
					return 0
				}

				path[i][j] = 1
				continue
			}

			if grid[i - 1][j - 1] != 1 {
				path[i][j] = path[i - 1][j] + path[i][j - 1]
			}
		}
	}

	return path[m][n]
}
