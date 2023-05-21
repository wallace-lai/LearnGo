// 一维动态规划

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/coin-change/?envType=study-plan-v2&envId=top-interview-150

// v1 : 36ms
// func depthFirstSearch(coins *[]int, f *[]int, amount int) int {
// 	if amount < 0 {
// 		return -1
// 	}
// 	if amount == 0 {
// 		return 0
// 	}
// 	if (*f)[amount - 1] != 0 {
// 		return (*f)[amount - 1]
// 	}

// 	result := math.MaxInt32
// 	for i := 0; i < len(*coins); i++ {
// 		ret := depthFirstSearch(coins, f, amount - (*coins)[i])
// 		if ret >= 0 && ret < result {
// 			result = ret + 1
// 		}
// 	}
// 	if result == math.MaxInt32 {
// 		result = -1
// 	}

// 	(*f)[amount - 1] = result
// 	return result
// }

// func coinChange(coins []int, amount int) int {
// 	if amount < 1 {
// 		return 0
// 	}

// 	f := make([]int, amount, amount)
// 	return depthFirstSearch(&coins, &f, amount)
// }

// v2 : 12ms
func getMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	f := make([]int, amount + 1, amount + 1)
	for i := 0; i < amount + 1; i++ {
		f[i] = math.MaxInt32
	}

	f[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				f[i] = getMin(f[i], f[i - coins[j]] + 1)
			}
		}
	}

	if f[amount] == math.MaxInt32 {
		return -1
	}
	return f[amount]
}

func main() {
	coins := []int {
		1, 2, 5,
	}

	fmt.Println(coinChange(coins, 11))
}