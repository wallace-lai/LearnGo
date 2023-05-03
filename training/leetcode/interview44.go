// 哈希表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/two-sum/?envType=study-plan-v2&id=top-interview-150

func twoSum(nums []int, target int) []int {
	result := make([]int, 2, 2)
	mapData := make(map[int]int, 0)	// value -> index
	for idx, num := range nums {
		v, ok := mapData[target - num]
		if ok == true {
			result[0] = v
			result[1] = idx
			break
		}

		mapData[num] = idx
	}

	return result
}

func main() {
	nums := []int {2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}