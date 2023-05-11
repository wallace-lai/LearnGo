// 二叉树层次遍历

package main

import (
	"fmt"
)

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64 {}
	}

	result := make([]float64, 0)

	idx := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) - idx > 0 {
		avg := float64(0)
		num := len(queue) - idx
		for i := 0; i < num; i++ {
			curr := queue[idx]
			idx++
			avg += float64(curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		avg = avg / float64(num)
		result = append(result, avg)
	}

	return result
}

func main() {

}