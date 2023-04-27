package main

import (
	"fmt"
)

// https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&id=top-interview-150

func removeElement(nums []int, val int) int {
	p := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[p] = nums[i]
			p++
		}
	}

	return p
}

/*
（一）
Go语言中for语句格式如下：
for init; condition; post {
	// do something
}
其中init语句中不允许像C语言那样使用逗号分隔多个赋值表达式，即：
for p = 0, i = 0; i < len(nums); i++ {}		// Go语言，非法
for (p = 0, i = 0; i < len(nums); i++) {}	// C语言，合法

（二）
Go语言中只支持后缀自增运算，不支持前缀自增运算，自减运算符亦然。即：
i++		// 合法
i--		// 合法
++i		// 非法
--i		// 非法

（三）
Go语言中的自增或者自减运算语句只能作为表达式使用，不能赋值给其他变量使用，这一点和C语言不一样。即：
nums[p++] = value		// Go语言，非法

nums[p] = value			// Go语言，合法
p++

nums[p++] = value		// C语言，合法
*/


func main() {
	nums := []int {3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}