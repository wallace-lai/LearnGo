package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/3sum/?envType=study-plan-v2&id=top-interview-150

// func threeSum(nums []int) [][]int {
// 	var result [][]int
// 	tuple := make([]int, 3)

// 	sort.Ints(nums)
// 	fmt.Println(nums)
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if j == i {
// 				continue
// 			}

// 			for k := 0; k < len(nums); k++ {
// 				if k == j || k == i {
// 					continue
// 				}

// 				if nums[i] + nums[j] + nums[k] == 0 {
// 					tuple[0] = nums[i]
// 					tuple[1] = nums[j]
// 					tuple[2] = nums[k]
// 					result = append(result, tuple)
// 				}
// 			}
// 		}
// 	}

// 	return result
// }

// func main() {
// 	nums := []int {-1, 0, 1, 2, -1, -4}
// 	fmt.Println(threeSum(nums))
// }