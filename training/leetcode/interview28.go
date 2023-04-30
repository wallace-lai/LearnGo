package main

import (
	"fmt"
)

// https://leetcode.cn/problems/container-with-most-water/?envType=study-plan-v2&id=top-interview-150

func getMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	result := 0
	i := 0
	j := len(height) - 1

	for i < j {
		if height[i] < height[j] {
			result = getMax(result, height[i] * (j - i))
			i++
		} else {
			result = getMax(result, height[j] * (j - i))
			j--
		}
	}

	return result
}

func main() {
	nums := []int {1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}