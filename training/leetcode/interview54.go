// 栈

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/min-stack/?envType=study-plan-v2&id=top-interview-150

type MinStack struct {
	stkLen1	int
	Stack1	[]int
	stkLen2	int
	Stack2	[]int
}


func Constructor() MinStack {
	minStk := MinStack{}
	minStk.stkLen1 = 0
	minStk.stkLen2 = 0
	return minStk
}


func (this *MinStack) Push(val int)  {
	this.Stack1 = append(this.Stack1, val)
	this.stkLen1 = len(this.Stack1)

	if this.stkLen2 == 0 || val < this.Stack2[this.stkLen2 - 1] {
		this.Stack2 = append(this.Stack2, val)
		this.stkLen2 = len(this.Stack2)
	} else {
		this.Stack2 = append(this.Stack2, this.Stack2[this.stkLen2 - 1])
		this.stkLen2 = len(this.Stack2)
	}
}


func (this *MinStack) Pop()  {
	this.Stack1 = this.Stack1[0:this.stkLen1 - 1]
	this.stkLen1 = len(this.Stack1)
	this.Stack2 = this.Stack2[0:this.stkLen2 - 1]
	this.stkLen2 = len(this.Stack2)
}


func (this *MinStack) Top() int {
	return this.Stack1[this.stkLen1 - 1]
}


func (this *MinStack) GetMin() int {
	return this.Stack2[this.stkLen2 - 1]
}

