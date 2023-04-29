package main

import (
	"fmt"
)

// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&id=top-interview-150

func revert(nums []int, beg int, end int) {
	for beg < end {
		tmp := nums[beg]
		nums[beg] = nums[end]
		nums[end] = tmp
		beg++
		end--
	}
}

// 24ms
func rotate(nums []int, k int) {
	size := len(nums)
	k = k % size

	if k == 0 {
		return
	}

	revert(nums, 0, size - 1)
	revert(nums, 0, k - 1)
	revert(nums, k, size - 1)
}

func main() {
	nums := []int {1, 2, 3, 4, 5, 6, 7}

	// rotate(nums, 0)
	// rotate(nums, 14)
	// rotate(nums, 1)
	// rotate(nums, 2)
	rotate(nums, 3)
	// rotate(nums, 4)
	// rotate(nums, 5)
	// rotate(nums, 6)
	fmt.Println(nums)
}