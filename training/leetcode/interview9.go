package main

import (
	"fmt"
)

// https://leetcode.cn/problems/jump-game-ii/?envType=study-plan-v2&id=top-interview-150

func getMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func jump(nums []int) int {
	step := 0
	currMaxDist := 0
	lastMaxDist := 0
	
	for i := 0; i < len(nums) - 1; i++ {
		currMaxDist = getMax(currMaxDist, i + nums[i])
		if i == lastMaxDist {
			lastMaxDist = currMaxDist
			step++
		}
	}

	return step
}

func main() {
	nums := []int {2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}