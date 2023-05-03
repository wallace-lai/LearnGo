// 哈希表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&id=top-interview-150

// v1 : timeout
// func findRoot(mapData map[int]int, num int) int {
// 	for {
// 		v, ok := mapData[num]
// 		if ok == false {
// 			return num
// 		}
// 		num = v
// 	}
// }

// func longestConsecutive(nums []int) int {
// 	mapData := make(map[int]int)
// 	for _, num := range nums {
// 		mapData[num] = num + 1
// 	}

// 	result := 0
// 	for _, num := range nums {
// 		root := findRoot(mapData, num)
// 		if root - num > result {
// 			result = root - num
// 		}
// 	}
// 	return result
// }

// v2 : 72ms
func longestConsecutive(nums []int) int {
	mapData := make(map[int]int)
	for _, num := range nums {
		mapData[num] = 1
	}

	result := 0
	for num, _ := range mapData {
		if mapData[num-1] != 1 {
			curr := num
			for mapData[num+1] == 1 {
				num++
			}
			if num-curr+1 > result {
				result = num - curr + 1
			}
		}
	}

	return result
}

func main() {
	// nums := []int {100, 4, 200, 1, 3, 2}
	// nums := []int {0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
