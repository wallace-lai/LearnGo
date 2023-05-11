// 二叉树层次遍历

package main

import (
	"fmt"
)

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/binary-tree-level-order-traversal/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int {}
	}

	result := make([][]int, 0)

	idx := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) - idx > 0 {
		level := make([]int, 0)

		num := len(queue) - idx
		for i := 0; i < num; i++ {
			curr := queue[idx]
			idx++
			level = append(level, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		result = append(result, level)
	}

	return result
}

func main() {

}