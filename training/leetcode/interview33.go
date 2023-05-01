package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/minimum-window-substring/?envType=study-plan-v2&id=top-interview-150

// func isCovered(sMapData, tMapData map[byte]int) bool {
// 	for k, v := range tMapData {
// 		if sMapData[k] < v {
// 			return false
// 		}
// 	}
// 	return true
// }

// v1 : 164ms
// func minWindow(s string, t string) (result string) {
// 	sLen := len(s)
// 	minLen := math.MaxInt32
// 	left, right := 0, 0

// 	// 获取t中字符统计情况
// 	sMapData := make(map[byte]int)
// 	tMapData := make(map[byte]int)
// 	for _, char := range t {
// 		tMapData[byte(char)]++
// 	}

// 	for right < sLen {
// 		in := byte(s[right])
// 		right++
// 		sMapData[in]++

// 		for isCovered(sMapData, tMapData) {
// 			if right - left < minLen {
// 				minLen = right - left
// 				result = s[left : right]
// 			}
// 			out := byte(s[left])
// 			left++
// 			sMapData[out]--
// 		}
// 	}

// 	return result
// }

// v2 : 40ms 将map换成slice
// best version : 0ms
func isCovered(sMapData, tMapData []int) bool {
	for k, v := range tMapData {
		if v > 0 && sMapData[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) (result string) {
	sLen := len(s)
	minLen := math.MaxInt32
	left, right := 0, 0

	sMapData := make([]int, 256)
	tMapData := make([]int, 256)
	for _, char := range t {
		tMapData[byte(char)]++
	}

	for right < sLen {
		in := byte(s[right])
		right++
		sMapData[in]++

		for isCovered(sMapData, tMapData) {
			if right-left < minLen {
				minLen = right - left
				result = s[left:right]
			}
			out := byte(s[left])
			left++
			sMapData[out]--
		}
	}

	return
}

func main() {
	// fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
}
