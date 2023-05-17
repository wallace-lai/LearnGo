// kadane算法

package main

// https://leetcode.cn/problems/maximum-subarray/?envType=study-plan-v2&id=top-interview-150

// v1 : 96ms
func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	f := make([]int, len(nums), len(nums))
	f[0] = nums[0]

	result := f[0]
	for i := 1; i < len(nums); i++ {
		f[i] = getMax(f[i - 1] + nums[i], nums[i])
		if f[i] > result {
			result = f[i]
		}
	}

	return result
}

