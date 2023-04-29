package main

import (
	"fmt"
)

// https://leetcode.cn/problems/is-subsequence/?envType=study-plan-v2&id=top-interview-150

// 0ms
func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)

	si := 0
	ti := 0
	for si < sLen && ti < tLen {
		if s[si] == t[ti] {
			si++
		}
		ti++
	}

	if si == sLen {
		return true
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}