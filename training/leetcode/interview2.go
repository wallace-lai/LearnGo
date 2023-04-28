package main

import (
	"fmt"
)
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&id=top-interview-150

func removeDuplicates(nums []int) int {
	size := len(nums)

	if size == 0 {
		return 0
	}

	slow := 1
	for fast := 1; fast < size; fast++ {
		if nums[fast] != nums[fast - 1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func main() {
	nums := []int {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}