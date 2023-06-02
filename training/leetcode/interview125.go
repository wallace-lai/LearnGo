// 位运算

package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/add-binary/?envType=study-plan-v2&envId=top-interview-150

func addBinary(a string, b string) string {
	ia, _ := strconv.ParseInt(a, 2, 32)
	ib, _ := strconv.ParseInt(b, 2, 32)
	fmt.Println(ia, ib)

	x := int(ia)
	y := int(ib)
	result, carry := 0, 0
	for y != 0 {
		result = x ^ y
		carry = (x & y) << 1
		fmt.Printf("result = %d, carry = %d\n", result, carry)
		x, y = result, carry
	}

	return strconv.FormatInt(int64(x), 2)
}

func main() {
	fmt.Println(addBinary("101011", "101101"))
}