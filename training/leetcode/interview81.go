// 二叉树

package main

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/?envType=study-plan-v2&id=top-interview-150

// v1 : 16ms
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}