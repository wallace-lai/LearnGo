// 多维动态规划

package main

// https://leetcode.cn/problems/maximal-square/?envType=study-plan-v2&envId=top-interview-150
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// v1 : 4ms
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	maxLen := 0
	m := len(matrix)
	n := len(matrix[0])

	f := make([][]int, m, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					f[i][j] = 1
				} else {
					f[i][j] = getMin(f[i - 1][j - 1], getMin(f[i - 1][j], f[i][j - 1])) + 1
				}

				if f[i][j] > maxLen {
					maxLen = f[i][j]
				}
			}
		}
	}

	return maxLen * maxLen
}

