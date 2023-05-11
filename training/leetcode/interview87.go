// 二叉搜索树

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func inorderTraverse(root *TreeNode, k int) int {
	if root.Left != nil {
		inorderTraverse(root.Left, k)
	}

	k--
	if k == 0 {
		fmt.Println("root.Val =", root.Val)
		return root.Val
	}
	
	if root.Right != nil {
		inorderTraverse(root.Right, k)
	}

	// unreachable code
	return -1
}

func kthSmallest(root *TreeNode, k int) int {
	return inorderTraverse(root, k)
}