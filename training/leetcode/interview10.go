package main

import (
	"fmt"
)

// https://leetcode.cn/problems/h-index/?envType=study-plan-v2&id=top-interview-150

func hIndex(citations []int) int {
	n := len(citations)
	f := make([]int, n + 1)		// f[i]表示引用数等于i的论文数量
	for _, citcitation := range citations {
		if citcitation >= n {
			f[n]++
		} else {
			f[citcitation]++
		}
	}

	total := 0
	for h := n; h >= 0; h-- {
		total += f[h]
		if total >= h {
			return h
		}
	}

	return 0
}

func main() {
	nums := []int {3, 0, 6, 1, 5}
	// nums := []int {1, 3, 1}
	fmt.Println(hIndex(nums))
}