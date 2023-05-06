// 栈

package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func evalRPN(tokens []string) int {
	numsLen := 0
	nums := make([]int, len(tokens), len(tokens))

	for i := 0; i < len(tokens); i++ {
		if  tokens[i] == "+" ||
			tokens[i] == "-" ||
			tokens[i] == "*" ||
			tokens[i] == "/" {
			lhs := nums[numsLen - 2]
			rhs := nums[numsLen - 1]
			numsLen -= 2

			switch tokens[i] {
			case "+":
				nums[numsLen] = lhs + rhs
			case "-":
				nums[numsLen] = lhs - rhs
			case "*":
				nums[numsLen] = lhs * rhs
			case "/":
				nums[numsLen] = lhs / rhs
			default:
			}
			numsLen++
		} else {
			result, _ := strconv.ParseInt(tokens[i], 10, 64)
			nums[numsLen] = int(result)
			numsLen++
		}
	}

	return nums[0]
}

func main() {
	tokens := []string {
		"2", "1", "+", "3", "*",
	}
	fmt.Println(evalRPN(tokens))
}