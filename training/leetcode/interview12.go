// 数组

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&id=top-interview-150

// v1 : 16ms
func productExceptSelf(nums []int) []int {
	leftProd := make([]int, len(nums))
	rightProd := make([]int, len(nums))

	leftProd[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProd[i] = leftProd[i - 1] * nums[i - 1]
	}

	rightProd[len(nums) - 1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightProd[i] = rightProd[i + 1] * nums[i + 1]
	}
	
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = leftProd[i] * rightProd[i]
	}

	return result
}

func main() {
	// nums := []int {1, 2, 3, 4}
	nums := []int {-1, 1, 0, -3, 3}

	fmt.Println(productExceptSelf(nums))
}


