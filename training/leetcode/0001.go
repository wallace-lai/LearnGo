package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var result = make([] int, 2)

	len := len(nums)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i] + nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}

	return result[:]
}

func main() {
	// nums := [] int {2, 7, 11, 15}
	// nums := [] int {3, 2, 4}
	nums := [] int {3, 3}
	fmt.Println(twoSum(nums, 6))
}