// 二叉树

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func build(in, post []int, iBeg, iEnd, pBeg, pEnd int) *TreeNode{
	var root *TreeNode
	root = nil

	if iBeg <= iEnd {
		rootVal := post[pEnd]
		root = &TreeNode {rootVal, nil, nil}

		idx := iBeg
		for in[idx] != rootVal {
			idx++
		}
		leftSize := idx - iBeg

		root.Left = build(in, post, iBeg, idx - 1, pBeg, pBeg + leftSize - 1)
		root.Right = build(in, post, idx + 1, iEnd, pBeg + leftSize, pEnd)
	}

	return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder, 0, len(inorder) - 1, 0, len(postorder) - 1)
}

func main() {

}