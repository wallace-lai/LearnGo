// 数学

package main

import "fmt"

// https://leetcode.cn/problems/plus-one/?envType=study-plan-v2&envId=top-interview-150

// v1 : 0ms
func plusOne(digits []int) []int {
	n := len(digits)

	sum := 0
	carry := 1
	for i := n - 1; i >= 0; i-- {
		sum = digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}
	if carry > 0 {
		result := make([]int, n + 1, n + 1)
		result[0] = carry
		for i := 1; i <= n; i++ {
			result[i] = digits[i - 1]
		}
		return result
	}

	return digits
}

func main() {
	nums := []int {
		1, 2, 3,
	}

	fmt.Println(plusOne(nums))
}