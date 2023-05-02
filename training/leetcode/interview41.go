// 哈希表

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/word-pattern/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func wordPattern(t string, s string) bool {
	words := strings.Split(s, " ")
	t2s := make(map[string]string)
	s2t := make(map[string]string)

	if len(t) != len(words) {
		return false
	}

	for i := 0; i < len(t); i++ {
		tWord := strconv.Itoa(int(t[i]))
		sWord := words[i]

		_, ok1 := t2s[tWord]
		_, ok2 := s2t[sWord]
		if ok1 == false && ok2 == false {
			t2s[tWord] = sWord
			s2t[sWord] = tWord
		} else if ok1 == true && ok2 == true {
			if t2s[tWord] != sWord || s2t[sWord] != tWord {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	// fmt.Println(strconv.Itoa(65))
}