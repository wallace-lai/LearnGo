// 数组

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/candy/?envType=study-plan-v2&id=top-interview-150

// v1 : 12ms
func candy(nums []int) int {
	size := len(nums)
	left := make([]int, size)
	right := make([]int, size)

	if size < 2 {
		return 1
	}

	left[0] = 1
	for i := 1; i < size; i++ {
		if nums[i] > nums[i - 1] {
			left[i] = left[i - 1] + 1
		} else {
			left[i] = 1
		}
	}

	right[size - 1] = 1
	for i := size - 2; i >= 0; i-- {
		if nums[i] > nums[i + 1] {
			right[i] = right[i + 1] + 1
		} else {
			right[i] = 1
		}
	}

	result := 0
	for i := 0; i < size; i++ {
		if left[i] > right[i] {
			result += left[i]
		} else {
			result += right[i]
		}
	}
	return result
}

func main() {
	nums := []int {1, 0, 2}
	// nums := []int {1, 2, 2}
	// nums := []int {1, 3, 2, 2, 1}
	fmt.Println(candy(nums))
}