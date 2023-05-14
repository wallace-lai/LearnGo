// 字典树

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/design-add-and-search-words-data-structure/?envType=study-plan-v2&id=top-interview-150

// v1 : 436ms
type WordDictionary struct {
	isWord 		bool
	children 	[26]*WordDictionary
}


func Constructor() WordDictionary {
	head := WordDictionary{isWord: false, children: [26]*WordDictionary {}}
	return head
}


func (this *WordDictionary) AddWord(word string)  {
	p := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if (p.children[idx] == nil) {
			p.children[idx] = &WordDictionary{isWord: false, children: [26]*WordDictionary{}}
		}
		p = p.children[idx]
	}
	p.isWord = true
}

func depthFirstSearch(this *WordDictionary, word string, i int, result *bool) {
	if this == nil {
		return
	}

	if i == len(word) {
		if this.isWord == true {
			*result = true
		}
		return
	}

	if word[i] == '.' {
		for _, v := range this.children {
			depthFirstSearch(v, word, i + 1, result)
		}
	} else {
		idx := word[i] - 'a'
		depthFirstSearch(this.children[idx], word, i + 1, result)

	}
}


func (this *WordDictionary) Search(word string) bool {
	result := false
	depthFirstSearch(this, word, 0, &result)
	return result
}

