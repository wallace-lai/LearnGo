// 二叉树

package main

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func getSum(root *TreeNode, sum int, result *int) {
	sum = sum * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		*result = *result + sum
		return
	}

	if root.Left != nil {
		getSum(root.Left, sum, result)
	}
	if root.Right != nil {
		getSum(root.Right, sum, result)
	}
}

func sumNumbers(root *TreeNode) int {
	result := 0

	getSum(root, 0, &result)
	return result
}

func main() {

}