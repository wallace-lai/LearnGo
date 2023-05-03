// 哈希表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&id=top-interview-150

func encodeString(s string) uint64 {
	prime := []uint64 {
		3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103,
	}

	sum := uint64(1)
	for _, char := range s {
		sum *= prime[int(char - 'a')]
	}
	return sum
}

// v1 : 20ms
func groupAnagrams(strs []string) [][]string {
	mapData := make(map[uint64][]string)
	for _, str := range strs {
		prod := encodeString(str)
		fmt.Println(prod)
		_, ok := mapData[prod]
		if ok == false {
			mapData[prod] = make([]string, 0)
		}
		mapData[prod] = append(mapData[prod], str)
	}

	idx := 0
	result := make([][]string, len(mapData), len(mapData))
	for _, v := range mapData {
		result[idx] = v
		idx++
	}
	return result
}

func main() {
	// strs := []string {
	// 	"eat",
	// 	"tea",
	// 	"tan",
	// 	"ate",
	// 	"nat",
	// 	"bat",
	// }

	// strs := []string {
	// 	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	// 	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
	// }

	strs := []string {
		"sat",
		"sin",
		"ins",
	}

	var result [][]string
	result = groupAnagrams(strs)
	for _, v := range result {
		fmt.Println(v)
	}
}