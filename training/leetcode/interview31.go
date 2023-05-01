package main

import (
	"fmt"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?envType=study-plan-v2&id=top-interview-150

func lengthOfLongestSubstring(s string) int {
	result := 0
	left, right := 0, 0
	window := make([]int, 256)

	for right < len(s) {
		var in = int(s[right])
		right++
		window[in]++

		for window[in] > 1 {
			var out = int(s[left])
			left++
			window[out]--
		}

		if right - left > result {
			result = right - left
		}
	}

	return result
}

func main() {
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}