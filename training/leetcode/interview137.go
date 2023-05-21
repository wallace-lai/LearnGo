// 一维动态规划

package main

// https://leetcode.cn/problems/climbing-stairs/?envType=study-plan-v2&envId=top-interview-150

// v1 : 0ms
// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1 
// 	}

// 	f := make([]int, n + 1, n + 1)
// 	f[1] = 1
// 	f[2] = 2
// 	for i := 3; i <= n; i++ {
// 		f[i] = f[i - 1] + f[i - 2]
// 	}

// 	return f[n]
// }

// v2 : 0ms
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	a := 1
	b := 2
	c := 0
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}

