package main

import (
	"fmt"
)

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/?envType=study-plan-v2&id=top-interview-150

func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	i := 0
	j := len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			result[0] = i + 1
			result[1] = j + 1
			return result
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return result
}

func main() {
	// nums := []int {2, 7, 11, 15}
	// nums := []int{2, 3, 4}
	nums := []int {-1, 0}

	fmt.Println(twoSum(nums, -1))
}
