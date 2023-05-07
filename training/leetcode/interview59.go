// 链表

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-two-sorted-lists/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode {0, nil}
	tail := &dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			tail = tail.Next
			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = tail.Next
			l2 = l2.Next
		}
	}

	lefted := l1
	if l2 != nil {
		lefted = l2
	}
	tail.Next = lefted

	return dummy.Next
}

func main() {

}