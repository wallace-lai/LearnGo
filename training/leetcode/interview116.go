// 二分查找

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/find-peak-element/?envType=study-plan-v2&envId=top-interview-150

func getValue(nums *[]int, idx int) int {
	n := len(*nums)
	if idx == -1 || idx == n {
		return math.MinInt64
	}
	return (*nums)[idx]
}

// v1 : 0ms
// func findPeakElement(nums []int) int {
// 	idx := (len(nums) - 1) / 2
// 	fmt.Println("idx =", idx)
// 	for !(getValue(&nums, idx - 1) < getValue(&nums, idx) && getValue(&nums, idx) > getValue(&nums, idx + 1)) {
// 		if getValue(&nums, idx) < getValue(&nums, idx + 1) {
// 			idx++
// 		} else {
// 			idx--
// 		}

// 		fmt.Println("idx =", idx)
// 	}

// 	return idx
// }

// v2 : 4ms
func findPeakElement(nums []int) int {
	n := len(nums)

	left := 0
	right := n - 1
	result := -1
	for left <= right {
		mid := left + (right - left) / 2
		if getValue(&nums, mid) > getValue(&nums, mid - 1) && 
			getValue(&nums, mid) > getValue(&nums, mid + 1) {
			result = mid
			break
		}

		if getValue(&nums, mid) < getValue(&nums, mid + 1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func main() {
	nums := []int {
		// 1, 2, 3, 1,
		1, 2, 1, 3, 5, 6, 4,
	}

	fmt.Println(findPeakElement(nums))
}