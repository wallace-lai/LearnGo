package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/3sum/?envType=study-plan-v2&id=top-interview-150

// 不成功的版本
// func threeSum(nums []int) [][]int {
// 	var result [][]int
// 	tuple := make([]int, 3)

// 	sort.Ints(nums)
// 	fmt.Println(nums)
// 	for i := 0; i < len(nums); i++ {
// 		if i == 0 || nums[i] != nums[i - 1] {
// 			for j := i + 1; j < len(nums); j++ {
// 				if j == i + 1 || nums[j] != nums[j - 1] {
// 					for k := j + 1; k < len(nums); k++ {
// 						if k == j + 1 || nums[k] != nums[k - 1] {
// 							if nums[i] + nums[j] + nums[k] == 0 {
// 								tuple[0] = nums[i]
// 								tuple[1] = nums[j]
// 								tuple[2] = nums[k]
// 								result = append(result, tuple)
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return result
// }

// timeout
// func threeSum(nums []int) [][]int {
// 	var result [][]int

// 	sort.Ints(nums)
// 	for i := 0; i < len(nums); i++ {
// 		if i == 0 || nums[i] != nums[i - 1] {
// 			for j := i + 1; j < len(nums); j++ {
// 				if j == i + 1 || nums[j] != nums[j - 1] {
// 					for k := j + 1; k < len(nums); k++ {
// 						if k == j + 1 || nums[k] != nums[k - 1] {
// 							if nums[i] + nums[j] + nums[k] == 0 {
// 								result = append(result, []int {nums[i], nums[j], nums[k]})
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return result
// }

func threeSum(nums []int) [][]int {
	var result [][]int

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		k := len(nums) - 1
		target := -1 * nums[i]
		for j := i + 1; j < len(nums); j++ {
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			for j < k && nums[j] + nums[k] > target {
				k--
			}

			if j == k {
				break
			}
			if nums[j] + nums[k] == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return result
}

func main() {
	nums := []int {-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}