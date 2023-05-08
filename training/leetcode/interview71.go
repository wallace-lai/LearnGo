// 二叉树

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/symmetric-tree/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func checkSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val &&
		   checkSymmetric(left.Left, right.Right) &&
		   checkSymmetric(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root, root)
}

func main() {

}