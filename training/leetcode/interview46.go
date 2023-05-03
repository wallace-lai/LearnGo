// 哈希表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/contains-duplicate-ii/?envType=study-plan-v2&id=top-interview-150

// v1 : 124ms, 54 case
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	mapData := make(map[int]int)	// num -> idx

// 	for idx, num := range nums {
// 		v, ok := mapData[num]
// 		if ok == true && idx - v <= k {
// 			return true
// 		}
// 		mapData[num] = idx
// 	}

// 	return false
// }

// v2 : 100ms
func containsNearbyDuplicate(nums []int, k int) bool {
	mapData := make(map[int]int)

	for idx, num := range nums {
		if idx > k {
			delete(mapData, nums[idx - k - 1])
		}

		_, ok := mapData[num]
		if ok == true {
			return true
		}
		mapData[num] = 1
	}

	return false
}

func main() {
	nums := []int {1, 2, 3, 1}
	fmt.Println(containsNearbyDuplicate(nums, 3))
}

