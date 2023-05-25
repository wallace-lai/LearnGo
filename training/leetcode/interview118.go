// 二分查找

package main

import (
	"fmt"
	"testing/quick"
)

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/?envType=study-plan-v2&envId=top-interview-150

// v1 : 8ms
func leftBound(nums *[]int, target int) int {
	n := len(*nums)
	if n == 0 {
		return -1
	}

	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right - left) / 2
		if (*nums)[mid] == target {
			right = mid - 1
		} else if (*nums)[mid] < target {
			left = mid + 1
		} else if (*nums)[mid] > target {
			right = mid - 1
		}
	}

	// 若target比数组中所有值都大
	if left == n {
		return -1
	}

	// 若target比数组中所有值都小或者target不存在于数组中
	if (*nums)[left] != target {
		return -1
	}

	return left
}

func rightBound(nums *[]int, target int) int {
	n := len(*nums)
	if n == 0 {
		return -1
	}

	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right - left) / 2
		if (*nums)[mid] == target {
			left = mid + 1
		} else if (*nums)[mid] < target {
			left = mid + 1
		} else if (*nums)[mid] > target {
			right = mid - 1
		}
	}

	// 若target比数组中所有值都小，必须先判断这个否则left - 1存在越界可能
	if right == -1 {
		return -1
	}

	// 若target比数组中所有值都大或者target不存在于数组中
	if (*nums)[left - 1] != target {
		return -1
	}

	return left - 1
}

func searchRange(nums []int, target int) []int {
	result := []int {-1, -1}
	result[0] = leftBound(&nums, target)
	result[1] = rightBound(&nums, target)
	return result
}


func main() {
	nums := []int {
		// 5, 7, 7, 8, 8, 10,
		1,
	}

	fmt.Println(searchRange(nums, 0))
}