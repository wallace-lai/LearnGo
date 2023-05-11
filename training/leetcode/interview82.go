// 二叉树层次遍历

package main

import "fmt"

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/binary-tree-right-side-view/?envType=study-plan-v2&id=top-interview-150

// v1 : 4ms
// func depthFirstSearch(root *TreeNode, depth int, mapData *map[int]int) {
// 	if _, ok := (*mapData)[depth]; !ok {
// 		(*mapData)[depth] = root.Val
// 	}

// 	if root.Right != nil {
// 		depthFirstSearch(root.Right, depth + 1, mapData)
// 	}
// 	if root.Left != nil {
// 		depthFirstSearch(root.Left, depth + 1, mapData)
// 	}
// }


// func rightSideView(root *TreeNode) []int {
// 	if root == nil {
// 		return []int {}
// 	}

// 	mapData := make(map[int]int)
// 	depthFirstSearch(root, 0, &mapData)
	
// 	result := make([]int, 0)
// 	for i := 0; i < len(mapData); i++ {
// 		result = append(result, mapData[i])
// 	}

// 	return result
// }

// v2 : 4ms
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int {}
	}
	
	result := make([]int, 0)

	idx := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) - idx > 0 {
		num := len(queue) - idx
		for i := 0; i < num; i++ {
			curr := queue[idx]
			idx++

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}

			if i == num - 1 {
				result = append(result, curr.Val)
			}
		}
	}

	return result
}

func main() {
	tree := []TreeNode {
		{1, nil, nil}, {2, nil, nil}, {3, nil, nil}, {5, nil, nil}, {4, nil, nil},
	}

	tree[0].Left = &tree[1]
	tree[0].Right = &tree[2]
	tree[1].Right = &tree[3]
	tree[2].Right = &tree[4]

	result := rightSideView(&tree[0])
	fmt.Println(result)
}