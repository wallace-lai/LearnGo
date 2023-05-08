// 二叉树

package main

// https://leetcode.cn/problems/same-tree/?envType=study-plan-v2&id=top-interview-150

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// v1 : 0ms
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		return p.Val == q.Val &&
			   isSameTree(p.Left, q.Left) &&
			   isSameTree(p.Right, q.Right)
	}

	return false
}

func main() {

}