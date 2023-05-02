// 哈希表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/isomorphic-strings/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
// func isIsomorphic(s string, t string) bool {
// 	st := make([]int, 256)
// 	ts := make([]int, 256)

// 	for i := 0; i < len(s); i++ {
// 		si := int(s[i])
// 		ti := int(t[i])
// 		if st[si] == 0 && ts[ti] == 0 {
// 			st[si] = ti
// 			ts[ti] = si
// 		} else if st[si] != 0 && ts[ti] != 0 {
// 			if st[si] != ti || ts[ti] != si {
// 				return false
// 			}
// 		} else if st[si] != 0 {
// 			if st[si] != ti {
// 				return false
// 			}
// 		} else if ts[ti] != 0 {
// 			if ts[ti] != si {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// v2 : 0ms
func isIsomorphic(s string, t string) bool {
	st := make([]int, 256)
	ts := make([]int, 256)

	for i := 0; i < len(s); i++ {
		si := int(s[i])
		ti := int(t[i])
		if st[si] == 0 && ts[ti] == 0 {
			st[si] = ti
			ts[ti] = si
		} else if st[si] != 0 && ts[ti] != 0 {
			if st[si] != ti || ts[ti] != si {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}