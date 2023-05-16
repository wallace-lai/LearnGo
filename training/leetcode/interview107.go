// 回溯
package main

import (
	"fmt"
)

// https://leetcode.cn/problems/word-search/?envType=study-plan-v2&id=top-interview-150

// v1 : 136ms
func checkExist(board *[][]byte, i, j int, vis *[][]bool, word string, idx int, result *bool) {
	if idx == len(word) - 1 {
		if (*board)[i][j] == byte(word[idx]) {
			// only update result if word exist
			*result = true
		}
		return
	}

	m := len(*board)
	n := len((*board)[0])
	di := []int {-1, 1, 0, 0}
	dj := []int {0, 0, -1, 1}
	if (*board)[i][j] == byte(word[idx]) {
		(*vis)[i][j] = true
		for k := 0; k < 4; k++ {
			newi := i + di[k]
			newj := j + dj[k]
			if (newi >= 0 && newi < m && newj >= 0 && newj < n && (*vis)[newi][newj] != true) {
				checkExist(board, newi, newj, vis, word, idx + 1, result)
			}
		}
		(*vis)[i][j] = false
	}
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	if m == 0 {
		return false
	}

	vis := make([][]bool, m, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n, n)
		for j := 0; j < n; j++ {
			vis[i][j] = false
		}
	}

	result := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			checkExist(&board, i, j, &vis, word, 0, &result)
			if result == true {
				return result
			}
		}
	}

	return result
}

func main() {
	board := [][]byte {
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(board, "ABCCED"))
}