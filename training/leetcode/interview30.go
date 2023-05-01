package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/minimum-size-subarray-sum/?envType=study-plan-v2&id=top-interview-150

func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt32
	left, right, sum := 0, 0, 0

	for right < len(nums) {
		// 当滑动窗口内数字之和小于target时扩展窗口
		if sum < target {
			sum += nums[right]
			right++
		}

		// 当滑动窗口内数字之和大于等于target时收缩窗口
		for sum >= target {
			if right - left < result {
				// 收缩前更新答案
				result = right - left
			}
			sum -= nums[left]
			left++
		}
	}

	if result == math.MaxInt32 {
		return 0
	}
	return result
}

func main() {
	nums := []int {2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}