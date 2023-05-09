// 二叉树

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/?envType=study-plan-v2&id=top-interview-150
func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// v1 : wrong answer 85/95
// func getPathSum(root *TreeNode, result *int) int {
// 	if root == nil {
// 		return 0
// 	}

// 	leftMax := getPathSum(root.Left, result)
// 	rightMax := getPathSum(root.Right, result)
// 	// root做为单独子树时，单独更新result
// 	if leftMax + rightMax + root.Val > *result {
// 		*result = leftMax + rightMax + root.Val
// 	}
	
// 	sum1 := root.Val								// 只取root节点时的最大路径和
// 	sum2 := root.Val + getMax(leftMax, rightMax)	// 取root及其子树节点时的最大路径和
// 	return getMax(sum1, sum2)
// }

// func maxPathSum(root *TreeNode) int {
// 	result := math.MinInt32

// 	ret := getPathSum(root, &result)
// 	return getMax(ret, result)
// }

// v2 : 16ms
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    var maxGain func(*TreeNode) int
    maxGain = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        // 递归计算左右子节点的最大贡献值
        // 只有在最大贡献值大于 0 时，才会选取对应子节点
        leftGain := max(maxGain(node.Left), 0)
        rightGain := max(maxGain(node.Right), 0)

        // 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
        priceNewPath := node.Val + leftGain + rightGain

        // 更新答案
        maxSum = max(maxSum, priceNewPath)

        // 返回节点的最大贡献值
        return node.Val + max(leftGain, rightGain)
    }
    maxGain(root)
    return maxSum
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func main() {
	fmt.Println(math.MinInt32)
}