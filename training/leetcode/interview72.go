// 二叉树

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
func build(pre, in []int, pBeg, pEnd, iBeg, iEnd int) *TreeNode {
	var root *TreeNode
	root = nil

	if pBeg <= pEnd {
		rootValue := pre[pBeg]
		root = &TreeNode { rootValue, nil, nil}

		idx := iBeg
		for in[idx] != rootValue {
			idx++
		}
		leftSize := idx - iBeg

		root.Left = build(pre, in, pBeg + 1, pBeg + leftSize, iBeg, idx - 1)
		root.Right = build(pre, in, pBeg + leftSize + 1, pEnd, idx + 1, iEnd)
	}

	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, len(preorder) - 1, 0, len(inorder) - 1)
}

func main() {

}