// 回溯
package main

import (
	"fmt"
)

// https://leetcode.cn/problems/n-queens-ii/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func checkValid(n int, vis *[][]bool, row int, col int) bool {
	for i := 0; i < row; i++ {
		if (*vis)[i][col] == true {
			return false
		}
	}
	for i, j := row - 1, col + 1; i >= 0 && j < n; i, j = i - 1, j + 1 {
		if (*vis)[i][j] == true {
			return false
		}
	}
	for i, j := row - 1, col - 1; i >= 0 && j >= 0; i, j = i - 1, j - 1 {
		if (*vis)[i][j] == true {
			return false
		}
	}

	return true
}

func depthFirstSearch(n int, row int, vis *[][]bool, track *[]int, result *[][]int) {
	if row == n {
		answer := make([]int, n, n)
		copy(answer, *track)
		*result = append(*result, answer)
		return
	}

	for i := 0; i < n; i++ {
		if checkValid(n, vis, row, i) {
			(*vis)[row][i] = true
			(*track) = append((*track), i)
			depthFirstSearch(n, row + 1, vis, track, result)
			*track = (*track)[:len(*track) - 1]
			(*vis)[row][i] = false
		}
	}
}

func totalNQueens(n int) int {
	result := [][]int {}
	track := []int {}
	vis := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			vis[i][j] = false
		}
	}

	depthFirstSearch(n, 0, &vis, &track, &result)
	return len(result)
}

func main() {
	fmt.Println(totalNQueens(5))
}