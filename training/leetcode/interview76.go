// 二叉树

package main

import "fmt"

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/path-sum/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
// func checkPathSum(root *TreeNode, target int) bool {
// 	if root.Left == nil && root.Right == nil && root.Val == target {
// 		return true
// 	}

// 	left, right := false, false
// 	if root.Left != nil {
// 		left = checkPathSum(root.Left, target - root.Val)
// 	}
// 	if root.Right != nil {
// 		right = checkPathSum(root.Right, target - root.Val)
// 	}
	
// 	return left || right
// }

// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	return checkPathSum(root, targetSum)
// }

// v2 : 4ms
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil{
        return targetSum == root.Val
    }

    return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}

func main() {

}