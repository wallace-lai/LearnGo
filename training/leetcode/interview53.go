// 栈

package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/simplify-path/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func simplifyPath(path string) string {
	size := 0
	stack := make([]string, len(path), len(path))
	components := strings.Split(path, "/")

	for i := 0; i < len(components); i++ {
		if components[i] == ".." {
			if size > 0 {
				size--
			}
			continue
		}
		if components[i] != "." && components[i] != "" {
			stack[size] = components[i]
			size++
		}
	}

	if size == 0 {
		return "/"
	}

	result := ""
	for i := 0; i < size; i++ {
		result += ("/" + stack[i])
	}
	return result
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))

	// nums := []int {1, 2, 3}
	// nums = nums[0:2]
	// fmt.Println(nums)
}