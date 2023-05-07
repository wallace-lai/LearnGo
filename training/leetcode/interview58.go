// 链表

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode.cn/problems/add-two-numbers/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	carry := 0
	dummy := ListNode {0, nil}
	tail := &dummy

	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + carry
		carry = sum / 10
		sum = sum % 10

		tail.Next = &ListNode {sum, nil}
		tail = tail.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	lefted := l1
	if l2 != nil {
		lefted = l2
	}

	for lefted != nil {
		sum = lefted.Val + carry
		carry = sum / 10
		sum = sum % 10

		tail.Next = &ListNode {sum, nil}
		tail = tail.Next

		lefted = lefted.Next
	}

	if carry > 0 {
		tail.Next = &ListNode {carry, nil}
	}

	return dummy.Next
}

func main() {

}