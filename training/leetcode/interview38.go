// 矩阵

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/game-of-life/?envType=study-plan-v2&id=top-interview-150

func getLiveCellNum(board [][]int, m, n, x, y int) int {
	dx := []int { -1, 0, 1}
	dy := []int { -1, 0, 1}

	result := 0
	for i := 0; i < len(dx); i++ {
		for j := 0; j < len(dy); j++ {
			newx := x + dx[i]
			newy := y + dy[j]
			if newx == x && newy == y {
				continue
			}
			if newx < 0 || newx >= m || newy < 0 || newy >= n {
				continue
			}

			if board[newx][newy] == 1 || board[newx][newy] == 3 {
				result++
			}
		}
	}

	return result
}

// v1 : 0ms
func gameOfLife(board [][]int)  {
	m := len(board)
	n := len(board[0])

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			liveNum := getLiveCellNum(board, m, n, x, y)
			if board[x][y] == 1 && liveNum < 2 {
				board[x][y] = 3		// live -> dead
				continue
			}
			if board[x][y] == 1 && (liveNum == 2 || liveNum == 3) {
				continue			// still alive
			}
			if board[x][y] == 1 && liveNum > 3 {
				board[x][y] = 3		// live -> dead
				continue

			}
			if board[x][y] == 0 && liveNum == 3 {
				board[x][y] = 2		// dead -> live
				continue
			}
		}
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] == 2 {
				board[x][y] = 1		// make cell alive
			}
			if board[x][y] == 3 {
				board[x][y] = 0		// make cell dead
			}
		}
	}
}

func main() {
	board := [][]int {
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	gameOfLife(board)
	fmt.Println(board)
}