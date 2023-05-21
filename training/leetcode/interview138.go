// 一维动态规划

package main

// https://leetcode.cn/problems/house-robber/?envType=study-plan-v2&envId=top-interview-150

func getMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

// v1 : 0ms
// f[i]表示前i间房屋能偷窃的最高金额，则有
// f[i] = max(f[i - 2] + nums[i], f[i - 1])
// func rob(nums []int) int {
// 	n := len(nums)
// 	if n == 1 {
// 		return nums[0]
// 	}

// 	f := make([]int, n, n)

// 	f[0] = nums[0]
// 	f[1] = getMax(nums[0], nums[1])
// 	for i := 2; i < n; i++ {
// 		f[i] = getMax(f[i - 2] + nums[i], f[i - 1])
// 	}

// 	return f[n - 1]
// }

// v2 : 0ms
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	f0 := nums[0]
	f1 := getMax(nums[0], nums[1])
	f2 := 0
	if n == 2 {
		return f1
	}

	for i := 2; i < n; i++ {
		f2 = getMax(f0 + nums[i], f1)
		f0 = f1
		f1 = f2
	}

	return f2
}