// 数组

package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/roman-to-integer/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
func romanToInt(s string) int {
	mapData := map[string]int {
		"M"  : 1000,
		"CM" : 900,
		"D"  : 500,
		"CD" : 400,
		"C"  : 100,
		"XC" : 90,
		"L"  : 50,
		"XL" : 40,
		"X"  : 10,
		"IX" : 9,
		"V"  : 5,
		"IV" : 4,
		"I"  : 1,
	}
	vecData := []string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	idx := 0
	result := 0
	start := 0
	for idx < len(s) {
		for i := start; i < len(vecData); i++ {
			if strings.HasPrefix(s[idx:], vecData[i]) {
				result += mapData[vecData[i]]
				idx += len(vecData[i])
				start = i
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}