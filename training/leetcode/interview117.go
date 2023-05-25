// 二分查找

package main

import "fmt"

// https://leetcode.cn/problems/search-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-interview-150

// v1 : 0ms
func search(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums := []int {
		4, 5, 6, 7, 0, 1, 2,
	}

	fmt.Println(search(nums, 99))
}