package main

import (
	"fmt"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/?envType=study-plan-v2&id=top-interview-150

func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i - 1] {
			result += prices[i] - prices[i - 1]
		}
	}
	return result
}

func main() {
	// prices := []int {7, 1, 5, 3, 6, 4}
	prices := []int {1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
}