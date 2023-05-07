// 链表

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode.cn/problems/rotate-list/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func rotateRight(head *ListNode, k int) *ListNode {
	size := 0
	for p := head; p != nil; p = p.Next {
		size++
	}
	if size == 0 {
		return head
	}

	k = k % size
	dummy := ListNode {0, head}
	curr := &dummy
	for i := 0; i < (size - k); i++ {
		curr = curr.Next
	}
	tail := curr.Next
	curr.Next = nil

	tmp := dummy.Next
	dummy.Next = tail
	curr = &dummy
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = tmp

	return dummy.Next
}

func main() {
	lists := []ListNode {
		{1, nil}, {2, nil}, {3, nil}, {4, nil}, {5, nil},
	}

	lists[0].Next = &lists[1]
	lists[1].Next = &lists[2]
	lists[2].Next = &lists[3]
	lists[3].Next = &lists[4]

	for i := 0; i < len(lists); i++ {
		fmt.Printf("addr of list[%d] : %p\n", i, &lists[i])
	}

	p := rotateRight(&lists[0], 2)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}