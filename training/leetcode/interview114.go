// 二分查找

package main

import "fmt"

// https://leetcode.cn/problems/search-insert-position/?envType=study-plan-v2&envId=top-interview-150

// func binarySearch(nums []int, target int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return -1
// 	}

// 	// 搜索区间[left, right]
// 	left := 0
// 	right := n - 1

// 	// 结束条件是left = right + 1
// 	for left <= right {
// 		mid := left + (right - left) / 2
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] < target {
// 			left = mid + 1
// 		} else if nums[mid] > target {
// 			right = mid - 1
// 		}
// 	}

// 	return -1
// }

// func leftBound(nums []int, target int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return -1
// 	}

// 	left := 0
// 	right := n - 1
// 	for left <= right {
// 		mid := left + (right - left) / 2
// 		if nums[mid] == target {
// 			right = mid - 1	// 收缩右侧边界
// 		} else if nums[mid] < target {
// 			left = mid + 1	// 收缩左侧边界
// 		} else if nums[mid] > target {
// 			right = mid - 1	// 收缩右侧边界
// 		}
// 	}

// 	// 1. 若target比数组中所有值都大
// 	if left == n {
// 		return -1
// 	}

// 	// 2. 若target比数组中所有值都小或者target不存在于数组中
// 	if nums[left] != target {
// 		return -1
// 	}

// 	return left
// }

func rightBound(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 搜索区间[left, right]
	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			left = mid + 1	// 收缩左侧边界
		} else if nums[mid] < target {
			left = mid + 1	// 收缩左侧边界
		} else if nums[mid] > target {
			right = mid - 1	// 收缩右侧边界
		}
	}

	// 情况1.2
	if right == -1 {
		return -1
	}

	// 情况1.1和1.3
	if nums[right] != target {
		return -1
	}

	// 情况2
	return right
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if (nums[mid] == target) {
			return mid
		} else if (nums[mid] < target) {
			left = mid + 1
		} else if (nums[mid] > target) {
			right = mid - 1
		}
	}

	return left
}

func main() {
	a := 1
	b := float64(a) + 1.1
	fmt.Println(a, b)
}
