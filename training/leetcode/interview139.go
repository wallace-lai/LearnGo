// 一维动态规划

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/word-break/?envType=study-plan-v2&envId=top-interview-150

// v1 : 0ms
func wordBreak(s string, wordDict []string) bool {
	minLen := math.MaxInt32
	maxLen := math.MinInt32
	mapData := make(map[string]bool)
	for _, v := range wordDict {
		mapData[v] = true
		if len(v) < minLen {
			minLen = len(v)
		}
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}

	n := len(s)
	f := make([]bool, n + 1, n + 1)

	f[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if f[j] {
				sub := string(s[j : i])
				if _, ok := mapData[sub]; ok {
					f[i] = true
					break
				}
			}
		}
	}

	return f[n]
}

func main() {
	words := []string {
		"leet",
		"code",
	}

	fmt.Println(wordBreak("leetcode", words))
}