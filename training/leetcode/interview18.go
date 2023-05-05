// 数组

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/length-of-last-word/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
// func lengthOfLastWord(s string) int {
// 	s = strings.TrimSpace(s)
// 	reg := regexp.MustCompile("\\s+")
// 	s = reg.ReplaceAllString(s, " ")

// 	words := strings.Split(s, " ")
// 	return len(words[len(words) - 1])
// }

// v2 : 0ms
func lengthOfLastWord(s string) int {
	result := 0

	end := len(s) - 1
	for s[end] == ' ' {
		end--
	}

	for i := end; i >= 0 && s[i] != ' '; i-- {
		result++
	}

	return result
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	// fmt.Println(lengthOfLastWord("hello world"))
}