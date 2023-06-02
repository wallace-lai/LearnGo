package main

import (
	"fmt"
)

// https://leetcode.cn/problems/single-number/?envType=study-plan-v2&envId=top-interview-150

// v1 : 16ms
func singleNumber(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}