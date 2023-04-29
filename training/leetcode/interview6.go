package main

import (
	"fmt"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&id=top-interview-150

// timeout
// func maxProfit(prices []int) int {
// 	size := len(prices)

// 	result := 0
// 	for i := 0; i < size; i++ {
// 		for j := i + 1; j < size; j++ {
// 			if prices[j] - prices[i] > result {
// 				result = prices[j] - prices[i]
// 			}
// 		}
// 	}

// 	return result
// }

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}

// 定义f[i]表示第i天能获得的最大收益
func maxProfit(prices []int) int {
	size := len(prices)
	if size < 1 {
		return 0
	}

	f := make([]int, size)
	lowest := prices[0]
	for i := 1; i < size; i++ {
		// 股票在第i天要么卖出，要么继续持有。因此最大收益是二者的最大值
		// （1）若股票在第i天卖出，则其收益为第i天的价格减去前i - 1天的股票最低价格
		// （2）若股票在第i天继续持有，则其收益和前i - 1天的最大收益一致
		f[i] = max(f[i - 1], prices[i] - lowest)

		if prices[i] < lowest {
			lowest = prices[i]
		}
	}

	return f[size - 1]
}

func main() {
	prices := []int {7, 1, 5, 3, 6}
	// prices := []int {7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices))
}