package main

import (
	"fmt"
)

// https://leetcode.cn/problems/jump-game/?envType=study-plan-v2&id=top-interview-150

func getMax(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}

func canJump(nums []int) bool {
	// maxDist表示从起点开始能达到的最远距离
	maxDist := 0
	for i := 0; i < len(nums); i++ {
		if i <= maxDist {
			maxDist = getMax(maxDist, i + nums[i])
			if maxDist >= len(nums) - 1 {
				return true
			}
		}
	}

	return maxDist >= len(nums) - 1
}

func main() {
	// nums := []int {2, 3, 1, 1, 4}
	nums := []int {3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}