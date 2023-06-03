package main


// https://leetcode.cn/problems/number-of-1-bits/?envType=study-plan-v2&envId=top-interview-150

// v1 : 4ms
func hammingWeight(num uint32) int {
	result := 0
	for num > 0 {
		num = num & (num - 1)
		result++
	}

	return result
}