// 图

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/surrounded-regions/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
func depthFirstSearch(board *[][]byte, m, n, i, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || (*board)[i][j] != 'O' {
		return
	}

	(*board)[i][j] = '1'

	depthFirstSearch(board, m, n, i - 1, j)
	depthFirstSearch(board, m, n, i + 1, j)
	depthFirstSearch(board, m, n, i, j - 1)
	depthFirstSearch(board, m, n, i, j + 1)
}

func solve(board [][]byte)  {
	m := len(board)
	n := len(board[0])

	// 从边缘为O处开始感染
	for i := 0; i < m; i++ {
		depthFirstSearch(&board, m, n, i, 0)
		depthFirstSearch(&board, m, n, i, n - 1)
	}
	for i := 0; i < n; i++ {
		depthFirstSearch(&board, m, n, 0, i)
		depthFirstSearch(&board, m, n, m - 1, i)
	}

	// 将未被感染的0设置为x，将已感染成1的恢复成0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}


func main() {

}