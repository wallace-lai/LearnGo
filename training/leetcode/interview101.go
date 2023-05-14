// 回溯
package main

import (
	"fmt"
)

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func depthFirstSearch(mapData *map[int]string, digit string, pos int, word string, result *[]string) {
	if pos == len(digit) {
		*result = append(*result, word)
		return
	}

	num := digit[pos] - '0'
	for _, char := range (*mapData)[int(num)] {
		depthFirstSearch(mapData, digit, pos + 1, word + string(char), result)
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	mapData := make(map[int]string)
	mapData[2] = "abc"
	mapData[3] = "def"
	mapData[4] = "ghi"
	mapData[5] = "jkl"
	mapData[6] = "mno"
	mapData[7] = "pqrs"
	mapData[8] = "tuv"
	mapData[9] = "wxyz"

	result := []string {}
	depthFirstSearch(&mapData, digits, 0, "", &result)
	return result
}

func main() {

}