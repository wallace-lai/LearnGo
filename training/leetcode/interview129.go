package main

import (
	"fmt"
)

// https://leetcode.cn/problems/single-number-ii/?envType=study-plan-v2&envId=top-interview-150

// v1 : 4ms
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < 32; i++ {
		sum := 0
		for _, num := range nums {
			sum += (num >> i) & 1
		}

		if sum % 3 > 0 {
			// 搞清楚为什么在Go里需要区分最高位
			if i == 31 {
				result = result - (1 << i)
			} else {
				result = result | (1 << i)
			}
		}
	}

	return result
}