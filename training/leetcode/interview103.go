// 回溯
package main

import (
	"fmt"
)

// https://leetcode.cn/problems/permutations/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func depthFirstSearch(nums *[]int, vis *[]bool, track *[]int, result *[][]int) {
	if len(*track) == len(*nums) {
		answer := make([]int, len(*track), len(*track))
		copy(answer, *track)
		*result = append(*result, answer)
		return
	}

	for i := 0; i < len(*nums); i++ {
		if (*vis)[i] == true {
			continue
		}
		*track = append(*track, (*nums)[i])
		(*vis)[i] = true
		depthFirstSearch(nums, vis, track, result)
		(*vis)[i] = false
		*track = (*track)[:len(*track) - 1]
	}
}

func permute(nums []int) [][]int {
	vis := make([]bool, len(nums), len(nums))
	for i := 0; i < len(vis); i++ {
		vis[i] = false
	}

	track := make([]int, 0)
	result := make([][]int, 0)
	depthFirstSearch(&nums, &vis, &track, &result)
	return result
}

func main() {
	nums := []int {1, 2, 3}
	fmt.Println(permute(nums))
}