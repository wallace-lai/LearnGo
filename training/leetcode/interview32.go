package main

import (
	"fmt"
)

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/?envType=study-plan-v2&id=top-interview-150

// v1 : 268ms
func findSubstring(s string, words []string) []int {
	var result []int
	length := len(s)
	wordNum := len(words)
	wordLen := len(words[0])

	for i := 0; i < length && i + wordNum * wordLen <= length; i++ {
		// 计算以i为起始位置的初始差异情况
		differ := make(map[string]int)
		for j := 0; j < wordNum; j++ {
			word := s[i + j * wordLen : i + (j + 1) * wordLen]
			differ[word]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}

		// 如果差异为0说明i是合法的起始位置
		if len(differ) == 0 {
			result = append(result, i)
		}
	}

	return result
}

// v2 : wrong answer ?
// func findSubstring(s string, words []string) []int {
// 	var result []int
// 	length := len(s)
// 	wordNum := len(words)
// 	wordLen := len(words[0])

// 	for i := 0; i < length && i + wordNum * wordLen <= length; i++ {
// 		differ := make(map[string]int)
// 		for j := 0; j < wordNum; j++ {
// 			word := s[i + j * wordLen : i + (j + 1) * wordLen]
// 			differ[word]++
// 		}
// 		for _, word := range words {
// 			differ[word]--
// 			if differ[word] == 0 {
// 				delete(differ, word)
// 			}
// 		}

// 		for start := i; start < length - wordNum * wordLen + 1; start += wordLen {
// 			if start != i {
// 				word := s[start + (wordNum - 1) * wordLen : start + wordNum * wordLen]
// 				differ[word]++
// 				if differ[word] == 0 {
// 					delete(differ, word)
// 				}

// 				word = s[start - wordLen : start]
// 				differ[word]--
// 				if differ[word] == 0 {
// 					delete(differ, word)
// 				}
// 			}
// 			if len(differ) == 0 {
// 				result = append(result, start)
// 			}
// 		}
// 	}

// 	return result
// }

func main() {
	// case 0
	// s := "barfoothefoobarman"
	// words := []string {
	// 	"foo",
	// 	"bar",
	// }

	// case 1
	// s := "wordgoodgoodgoodbestword"
	// words := []string {
	// 	"word",
	// 	"good",
	// 	"best",
	// 	"wrod",
	// }

	// case 2
	// s := "barfoofoobarthefoobarman"
	// words := []string {
	// 	"bar",
	// 	"foo",
	// 	"the",
	// }

	// case 3
	s := "wordgoodgoodgoodbestword"
	words := []string {
		"word",
		"good",
		"best",
		"good",
	}

	fmt.Println(findSubstring(s, words))
}