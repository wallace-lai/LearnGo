// 位运算

package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/add-binary/?envType=study-plan-v2&envId=top-interview-150

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// v1 : 4ms
func addBinary(a string, b string) string {
	result := ""
	carry := 0

	la := len(a)
	lb := len(b)
	n := getMax(la, lb)

	for i := 0; i < n; i++ {
		if i < la {
			carry += int(a[la - i - 1] - '0')
		}
		if i < lb {
			carry += int(b[lb - i - 1] - '0')
		}
		result = strconv.Itoa(carry % 2) + result
		carry /= 2
	}
	if carry > 0 {
		result = "1" + result
	}

	return result
}

func main() {
	fmt.Println(addBinary("101011", "101101"))
}