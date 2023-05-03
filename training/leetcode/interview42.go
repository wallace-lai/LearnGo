// 哈希表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/valid-anagram/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func isAnagram(s string, t string) bool {
	counter := make([]int, 26)

	for _, v := range s {
		counter[byte(v) - 'a']++
	}
	for _, v := range t {
		counter[byte(v) - 'a']--
	}
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	// fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}