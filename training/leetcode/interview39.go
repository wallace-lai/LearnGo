// 哈希表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/ransom-note/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func canConstruct(ransomNote string, magazine string) bool {
	count := make([]int, 26)
	for _, v := range ransomNote {
		count[byte(v) - 'a']++
	}

	for _, v := range magazine {
		count[byte(v) - 'a']--
	}

	for _, v := range count {
		if v > 0 {
			return false
		}
	}

	return true
}

func main() {

}