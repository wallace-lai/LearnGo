// 分治

package main

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/?envType=study-plan-v2&id=top-interview-150

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// v1 : 0ms
func convertToBST(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := TreeNode {nums[mid], nil, nil}
	root.Left = convertToBST(nums, left, mid - 1)
	root.Right = convertToBST(nums, mid + 1, right)
	return &root
}

func sortedArrayToBST(nums []int) *TreeNode {
	return convertToBST(nums, 0, len(nums) - 1)
}