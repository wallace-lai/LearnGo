// 二叉树

package main

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/?envType=study-plan-v2&id=top-interview-150

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// v1 : 4ms
// func getMax(a int, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return getMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
// }

// v2 : 0ms
type Queue struct {
	VecData 	[]*TreeNode
	Beg			int
	End			int
}

func Constructor() Queue {
	return Queue{VecData: make([]*TreeNode, 0), Beg: 0, End: 0 }
}

func (this *Queue) Push(node *TreeNode) {
	this.VecData = append(this.VecData, node)
	this.End++
}

func (this *Queue) Pop() *TreeNode {
	node := this.VecData[this.Beg]
	this.Beg++
	return node
}

func (this *Queue) Front() *TreeNode {
	return this.VecData[this.Beg]
}

func (this *Queue) Size() int {
	return this.End - this.Beg
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0;
	}

	q := Constructor()
	q.Push(root)
	depth := 0

	for q.Size() > 0 {
		num := q.Size()
		for i := 0; i < num; i++ {
			curr := q.Pop()
			if curr.Left != nil {
				q.Push(curr.Left)
			}
			if curr.Right != nil {
				q.Push(curr.Right)
			}
		}

		depth++
	}

	return depth
}

func main() {

}