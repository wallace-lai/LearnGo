// 二叉树

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/invert-binary-tree/?envType=study-plan-v2&id=top-interview-150

func invertRecursively(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = invertRecursively(root.Right)
	root.Right = invertRecursively(tmp)
	return root
}

// v1 : 0ms
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)
	return root
}

func main() {

}