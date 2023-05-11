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

// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/?envType=study-plan-v2&id=top-interview-150

// v1 : 12ms
func inorderTraverse(root *TreeNode, prev *int, result *int) {
	if root.Left != nil {
		inorderTraverse(root.Left, prev, result)
	}

	if *prev >= 0 {
		abs := 0
		if root.Val >= *prev {
			abs = root.Val - *prev
		} else {
			abs = *prev - root.Val
		}
		if abs < *result {
			*result = abs
		}
	}
	*prev = root.Val

	if root.Right != nil {
		inorderTraverse(root.Right, prev, result)
	}
}

func getMinimumDifference(root *TreeNode) int {
	prev := -1
	result := math.MaxInt32

	inorderTraverse(root, &prev, &result)
	return result
}
