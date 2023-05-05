// 区间

package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/summary-ranges/?envType=study-plan-v2&id=top-interview-150

// 0ms
func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return []string{}
	}

	beg := 0
	result := make([]string, 0, n)
	for i := 1; i < n; i++ {
		if nums[i] != nums[i - 1] + 1 {
			if i - beg == 1 {
				result = append(result, strconv.Itoa(nums[beg]))
			} else {
				result = append(result, strconv.Itoa(nums[beg]) + "->" + strconv.Itoa(nums[i - 1]))
			}

			beg = i
		}
	}
	if len(nums) - beg == 1{
		result = append(result, strconv.Itoa(nums[beg]))
	} else {
		result = append(result, strconv.Itoa(nums[beg]) + "->" + strconv.Itoa(nums[len(nums) - 1]))
	}
	return result
}

func main() {
	// nums := []int {0, 1, 2, 4, 5, 7}
	nums := []int{}
	fmt.Println(summaryRanges(nums))
}
