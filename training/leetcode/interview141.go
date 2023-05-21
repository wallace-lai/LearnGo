// 一维动态规划

package main

import "math"

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// v1 : 56ms
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	result := math.MinInt32
	f := make([]int, n, n)
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = getMax(f[i], f[j] + 1)
			}
		}

		if f[i] > result {
			result = f[i]
		}
	}

	return result
}

