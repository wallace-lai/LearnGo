// 二叉树

package main

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/count-complete-tree-nodes/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSum := countNodes(root.Left)
	rightSum := countNodes(root.Right)
	return leftSum + rightSum + 1
}

