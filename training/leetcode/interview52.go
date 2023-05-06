// 栈

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/valid-parentheses/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func isValid(s string) bool {
	size := 0
	stack := make([]byte, len(s), len(s))

	for _, char := range s {
		if size > 0 {
			if  stack[size - 1] == '(' && char == ')' ||
				stack[size - 1] == '{' && char == '}' ||
				stack[size - 1] == '[' && char == ']' {
				size--
				continue
			}
		}
		stack[size] = byte(char)
		size++
	}

	return size == 0
}

func main() {
	// fmt.Println(isValid("()"))
	// fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(])"))
}