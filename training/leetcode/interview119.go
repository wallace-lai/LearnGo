// 二分查找

package main

import "fmt"

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-interview-150

// v1 : 0ms
func findMin(nums []int) int {
	n := len(nums)
	
	left := 0
	right := n - 1
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[right]
}

func main() {
	nums := []int {
		// 3, 4, 5, 1, 2,
		4, 5, 6, 7, 0, 1, 2,
		// 11, 13, 15, 17,
	}

	fmt.Println(findMin(nums))
}