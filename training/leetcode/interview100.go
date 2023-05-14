// 字典树

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/word-search-ii/?envType=study-plan-v2&id=top-interview-150

type Trie struct {
	isWord 		bool
	children 	[26]*Trie
}

func Constructor() Trie {
	head := Trie {isWord: false, children: [26]*Trie {}}
	return head
}


func (this *Trie) Insert(word string)  {
	p := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if (p.children[idx] == nil) {
			p.children[idx] = &Trie{isWord: false, children: [26]*Trie {}}
		}
		p = p.children[idx]
	}
	p.isWord = true
}


func (this *Trie) Search(word string) bool {
	p := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if (p.children[idx] != nil) {
			p = p.children[idx]
		} else {
			return false
		}
	}
	return p.isWord
}


func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if (p.children[idx] != nil) {
			p = p.children[idx]
		} else {
			return false
		}
	}
	return true
}

func findWords(board [][]byte, words []string) []string {

}

