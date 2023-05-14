// 回溯
package main

import (
	"sort"
)

// https://leetcode.cn/problems/combination-sum/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func depthFirstSearch(nums *[]int, target int, track *[]int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		answer := make([]int, len(*track), len(*track))
		copy(answer, *track)
		*result = append(*result, answer)
		return
	}

	for i := 0; i < len(*nums); i++ {
		size := len(*track)
		if size > 0 && (*track)[size - 1] > (*nums)[i] {
			continue
		}

		*track = append(*track, (*nums)[i])
		depthFirstSearch(nums, target - (*nums)[i], track, result)
		*track = (*track)[:len(*track) - 1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	track := make([]int, 0)
	result := make([][]int, 0)
	depthFirstSearch(&candidates, target, &track, &result)
	return result
}

