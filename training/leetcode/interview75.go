// 二叉树

package main

import "fmt"

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func flat(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := flat(root.Left)
	right := flat(root.Right)
	if left == nil {
		root.Left = nil
		root.Right = right
		return root
	}

	prev := left
	curr := left.Right
	for curr != nil {
		prev = curr
		curr = curr.Right
	}
	prev.Right = right

	root.Left = nil
	root.Right = left
	return root
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	flat(root)
}

func dumpTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("root.Val = %d, root.Left = %p, root.Right = %p\n", root.Val, root.Left, root.Right)
	dumpTree(root.Left)
	dumpTree(root.Right)
}

func main() {
	// tree := []TreeNode {
	// 	{1, nil, nil}, {2, nil, nil}, {5, nil, nil}, {3, nil, nil}, {4, nil, nil}, {6, nil, nil},
	// }

	// tree[0].Left = &tree[1]
	// tree[0].Right = &tree[2]

	tree := []TreeNode {
		{1, nil, nil},
		{2, nil, nil},
		{5, nil, nil},
		{3, nil, nil},
		{4, nil, nil},
		{6, nil, nil},
	}
	tree[0].Left = &tree[1]
	tree[0].Right = &tree[2]
	tree[1].Left = &tree[3]
	tree[1].Right = &tree[4]
	tree[2].Right = &tree[5]

	for i := 0; i < len(tree); i++ {
		fmt.Printf("root.Val = %d, root = %p\n", tree[i].Val, &tree[i])
	}
	flatten(&tree[0])
	dumpTree(&tree[0])
}
