// 数组

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/integer-to-roman/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func intToRoman(num int) string {
	mapData := map[int]string {
		1000 : "M",
		900 : "CM",
		500 : "D",
		400 : "CD",
		100 : "C",
		90 : "XC",
		50 : "L",
		40 : "XL",
		10 : "X",
		9 : "IX",
		5 : "V",
		4 : "IV",
		1 : "I",
	}
	vecData := []int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	for num > 0 {
		for i := 0; i < len(vecData); i++ {
			if num >= vecData[i] {
				result += mapData[vecData[i]]
				num -= vecData[i]
				break
			}
		}
	}

	return result
}

func main() {
	// fmt.Println(intToRoman(3))
	// fmt.Println(intToRoman(4))
	// fmt.Println(intToRoman(9))
	// fmt.Println(intToRoman(58))
	// fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(20))
}