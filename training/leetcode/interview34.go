// 矩阵

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/valid-sudoku/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func isValidSudoku(board [][]byte) bool {
	var rows [9][10]int
	var cols [9][10]int
	var blks [9][10]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == byte('.') {
				continue
			}
			num := board[i][j] - '0'
			fmt.Println(num)
			if rows[i][num] == 1 || cols[j][num] == 1 || blks[i / 3 * 3 + j / 3][num] == 1 {
				return false
			}
			rows[i][num] = 1
			cols[j][num] = 1
			blks[i / 3 * 3 + j / 3][num] = 1
		}
	}

	return true
}

