package main

import (
	"fmt"
)

// https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&id=top-interview-150

// 12ms
// func majorityElement(nums []int) int {
// 	size := len(nums)
// 	sort.Ints(nums)
// 	return nums[size / 2]
// }

// 12ms
func majorityElement(nums []int) int {
	size := len(nums)
	if size < 2 {
		return nums[0]
	}

	majority := nums[0]
	count := 1
	for i := 1; i < size; i++ {
		if nums[i] != majority {
			count -= 1
		} else {
			count += 1
		}

		if count == 0 {
			majority = nums[i]
			count = 1
		}
	}

	return majority
}

/*
一、Go语言支持三元运算符
	count += (value == majority) ? 1 : -1	// Go语言，非法

二、使用sort包对切片或数组进行排序
	sort.Ints(nums)		// 对int型切片进行排序
	sort.Float64s(nums)	// 对float64型切片进行排序
*/

func main() {
	nums := []int {3, 2, 3}
	// nums := []int {2, 2, 1, 1, 1, 2, 2}
	// nums := []int {3, 2, 3}
	// nums := []int {6, 5, 5}
	fmt.Println(majorityElement(nums))
}