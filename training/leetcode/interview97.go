// 图的广度优先搜索

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/word-ladder/?envType=study-plan-v2&id=top-interview-150

// v1 : 396ms
func checkValid(curr, next string, vis *map[string]bool) bool {
	sum := 0
	for i := 0; i < len(curr); i++ {
		if curr[i] != next[i] {
			sum++
		}
	}
	if sum != 1 {
		return false
	}

	if _, ok := (*vis)[next]; ok {
		return false
	}

	return true
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 判断endWord是否在wordList当中
	found := false
	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			found = true
			break
		}
	}
	if !found {
		return 0
	}

	idx := 0
	queue := make([]string, 0)
	vis := make(map[string]bool)

	result := 0
	queue = append(queue, beginWord)
	vis[beginWord] = true

	for len(queue) - idx > 0 {
		result++

		size := len(queue) - idx
		for i := 0; i < size; i++ {
			curr := queue[idx]
			idx++
			if curr == endWord {
				return result
			}

			for _, next := range wordList {
				if checkValid(curr, next, &vis) {
					queue = append(queue, next)
					vis[next] = true
				}
			}
		}
	}

	return 0
}

// v2 : 双端BFS ?

func main() {
	beg := "hit"
	end := "cog"

	wordlist := []string {
		"hot",
		"dot",
		"dog",
		"lot",
		"log",
		"cog",
	}

	fmt.Println(ladderLength(beg, end, wordlist))
}