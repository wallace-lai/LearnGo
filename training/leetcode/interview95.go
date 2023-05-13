// 图的广度优先搜索

package main

import "fmt"

// https://leetcode.cn/problems/snakes-and-ladders/?envType=study-plan-v2&id=top-interview-150

// v1 : 12ms
func snakesAndLadders(board [][]int) int {
	n := len(board)

	vis := make([]bool, n * n + 1)
	for i := 0; i < len(vis); i++ {
		vis[i] = false
	}

	idx := 0
	queue := make([]int, 0)
	queue = append(queue, 1)
	vis[1] = true

	result := -1
	for len(queue) - idx > 0 {
		result++

		size := len(queue)  - idx
		for i := 0; i < size; i++ {
			curr := queue[idx]
			idx++

			if curr == n * n {
				return result
			}

			for x := 1; x <= 6; x++ {
				next := curr + x
				if next >= n * n {
					next = n * n
				}

				row := (next - 1) / n
				col := 0
				if row & 1 == 1 {
					col = n - 1 - ((next - 1) % n)
				} else {
					col = (next - 1) % n
				}
				row = n - 1 - row

				if board[row][col] > -1 {
					next = board[row][col]
				}

				if vis[next] == false {
					vis[next] = true
					queue = append(queue, next)
				}
			}
		}
	}

	return -1
}

// func convertToPosition(n int, num int) (row, col int) {
// 	col = 0
// 	row = (num - 1) / n
// 	if row & 1 == 1 {
// 		col = n - 1 - ((num - 1) % n)
// 	} else {
// 		col = (num - 1) % n
// 	}
// 	row = n - 1 - row

// 	return row, col
// }

func main() {
	// board := [][]int {
	// 	{-1,-1,-1,-1,-1,-1},
	// 	{-1,-1,-1,-1,-1,-1},
	// 	{-1,-1,-1,-1,-1,-1},
	// 	{-1,35,-1,-1,13,-1},
	// 	{-1,-1,-1,-1,-1,-1},
	// 	{-1,15,-1,-1,-1,-1},
	// }

	// board := [][]int {
	// 	{-1, -1},
	// 	{-1, 3},
	// }

	board := [][]int {
		{1, 1, -1},
		{1, 1, 1},
		{-1, 1, 1},
	}

	fmt.Println(snakesAndLadders(board))
}