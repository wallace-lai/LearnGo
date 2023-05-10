// 二叉树

package main

import (
	"fmt"
)

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/count-complete-tree-nodes/?envType=study-plan-v2&id=top-interview-150

// v1 : 20ms
// type BSTIterator struct {
// 	Idx 		int
// 	Size 		int
// 	VecData 	[]int
// }

// func inorderTraverse(root *TreeNode, vec *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	inorderTraverse(root.Left, vec)
// 	*vec = append(*vec, root.Val)
// 	inorderTraverse(root.Right, vec)
// }

// func Constructor(root *TreeNode) BSTIterator {
// 	iter := BSTIterator {Idx: 0, Size: 0, VecData: make([]int, 0)}
// 	vec := make([]int, 0)
// 	inorderTraverse(root, &vec)
// 	iter.Size = len(vec)
// 	iter.VecData = vec
// 	return iter
// }

// func (this *BSTIterator) Next() int {
// 	if this.Idx < this.Size {
// 		result := this.VecData[this.Idx]
// 		this.Idx++
// 		return result
// 	}

// 	// unreachable code
// 	return -1
// }


// func (this *BSTIterator) HasNext() bool {
// 	return this.Idx < this.Size
// }

// v2 : 20ms
type BSTIterator struct {
	curr	*TreeNode
	stack 	[]*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator {curr: root, stack: make([]*TreeNode, 0)}
	return iter
}


func (this *BSTIterator) Next() int {
	for this.curr != nil {
		this.stack = append(this.stack, this.curr)
		this.curr = this.curr.Left
	}

	this.curr = this.stack[len(this.stack) - 1]
	this.stack = this.stack[0 : len(this.stack) - 1]
	result := this.curr.Val
	this.curr = this.curr.Right
	return result
}


func (this *BSTIterator) HasNext() bool {
	return this.curr != nil || len(this.stack) > 0
}
