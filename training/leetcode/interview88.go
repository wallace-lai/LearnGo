// 二叉搜索树

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/validate-binary-search-tree/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
// func inorderTraverse(root *TreeNode, array *[]int) {
// 	if (root == nil) {
// 		return
// 	}

// 	inorderTraverse(root.Left, array)
// 	*array = append(*array, root.Val)
// 	inorderTraverse(root.Right, array)
// }

// func isValidBST(root *TreeNode) bool {
// 	vec := make([]int, 0)
// 	inorderTraverse(root, &vec)

// 	for i := 0; i < len(vec) - 1; i++ {
// 		if vec[i] >= vec[i + 1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// v2 : 8ms
func checkRangeValid(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return checkRangeValid(root.Left, lower, root.Val) &&
	       checkRangeValid(root.Right, root.Val, upper)
}

func isValidBST(root *TreeNode) bool {
	return checkRangeValid(root, math.MinInt64, math.MaxInt64)
}

func main() {

}