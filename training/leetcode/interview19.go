// 数组

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/longest-common-prefix/?envType=study-plan-v2&id=top-interview-150

func longestCommonPrefix(strs []string) string {
	num := len(strs)
	result := len(strs[0])

	for i := 1; i < num; i++ {
		idx := 0
		for idx < len(strs[i]) && idx < len(strs[0]) && strs[i][idx] == strs[0][idx] {
			idx++
		}
		if idx < result {
			result = idx
		}
	}

	return strs[0][0:result]
}

func main() {
	// strs := []string {
	// 	"flower",
	// 	"flow",
	// 	"flight",
	// }

	strs := []string {
		"dog",
		"racecar",
		"car",
	}

	fmt.Println(longestCommonPrefix(strs))
}