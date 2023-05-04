// 数组

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/gas-station/?envType=study-plan-v2&id=top-interview-150

// v1 : 104ms
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	idx := 0
	left := 0
	minLeft := math.MaxInt32

	for i := 0; i < n; i++ {
		left += gas[i] - cost[i]
		if left < minLeft {
			minLeft = left
			idx = i
		}
	}

	if left < 0 {
		return -1
	}

	return (idx + 1) % n
}

func main() {
	gas := []int {1, 2, 3, 4, 5}
	cost := []int {3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}