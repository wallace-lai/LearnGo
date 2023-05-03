// 哈希表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/happy-number/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func isHappy(n int) bool {
	sum := 0
	mapData := make(map[int]int)

	mapData[n] = 1
	for n != 1 {
		sum = 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum

		_, ok := mapData[sum]
		if ok == true {
			return false
		}
		mapData[sum] = 1
	}

	return true
}

func main() {
	fmt.Println(isHappy(19))
}