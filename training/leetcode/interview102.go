// 回溯
package main

import (
	"fmt"
)

// https://leetcode.cn/problems/combinations/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
func depthFirstSearch(n int, vis *[]bool, k int, pos int, track *[]int, result *[][]int) {
	if k == 0 {
		answer := make([]int, len(*track), len(*track))
		copy(answer, *track)
		*result = append(*result, answer)
		return
	}

	for i := pos; i <= n; i++ {
		if (*vis)[i] == true {
			continue
		}

		(*vis)[i] = true
		*track = append(*track, i)
		depthFirstSearch(n, vis, k - 1, i + 1, track, result)
		*track = (*track)[:len(*track) - 1]
		(*vis)[i] = false
	}
}

func combine(n int, k int) [][]int {
	vis := make([]bool, n + 1, n + 1)
	for i := 0; i < n + 1; i++ {
		vis[i] = false
	}

	track := []int {}
	result := [][]int {}
	depthFirstSearch(n, &vis, k, 1, &track, &result)
	return result
}

func main() {
	fmt.Println(combine(4, 2))
}