// 数学

package main

// https://leetcode.cn/problems/powx-n/?envType=study-plan-v2&envId=top-interview-150

// v1 : 0ms
func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	y := fastPow(x, n / 2)
	if n % 2 == 0 {
		return y * y
	}
	return y * y * x
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return fastPow(x, n)
	}
	return 1.0 / fastPow(x, -n)
}