// 链表

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode.cn/problems/reverse-linked-list-ii/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func reverseList(head *ListNode, tail *ListNode) *ListNode {
	dummy := ListNode {0, tail}

	for head != nil {
		p := head
		head = head.Next
		p.Next = dummy.Next
		dummy.Next = p
	}

	return dummy.Next
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := ListNode {0, head}

	var (
		pHead *ListNode = nil
		pTail *ListNode = nil

		pBeg *ListNode = nil
		pEnd *ListNode = nil
	)

	p := &dummy
	for i := 1; i <= right; i++ {
		if i == left {
			pHead = p
			pBeg = p.Next
		}
		if i == right {
			pEnd = p.Next
			pTail = pEnd.Next
		}

		p = p.Next
	}
	pEnd.Next = nil

	pHead.Next = reverseList(pBeg, pTail)
	return dummy.Next
}

func main() {
	lists := []ListNode {
		{1, nil}, {2, nil}, {3, nil}, {4, nil}, {5, nil},
	}

	// lists[0].Next = &lists[1]
	// lists[1].Next = &lists[2]
	// lists[2].Next = &lists[3]
	// lists[3].Next = &lists[4]

	for i := 0; i < len(lists); i++ {
		fmt.Printf("addr of list[%d] : %p\n", i, &lists[i])
	}

	p := reverseBetween(&lists[4], 1, 1)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}

}