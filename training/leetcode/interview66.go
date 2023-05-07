// 链表

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode.cn/problems/partition-list/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	result := ListNode {0, nil}
	toAdd := &result

	dummy := ListNode {0, head}
	prev := &dummy
	curr := dummy.Next
	for curr != nil {
		// 将小于x的节点从原链表中摘链并挂在result链中
		if curr.Val < x {
			prev.Next = curr.Next
			curr.Next = nil
			toAdd.Next = curr
			toAdd = toAdd.Next

			curr = prev.Next
			continue
		}

		prev = curr
		curr = curr.Next
	}
	// 将剩余大于等于x的节点挂到result链表末尾
	toAdd.Next = dummy.Next
	dummy.Next = nil
	
	return result.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	lists := []ListNode {
		{1, nil}, {4, nil}, {3, nil}, {2, nil}, {5, nil}, {2, nil},
	}

	lists[0].Next = &lists[1]
	lists[1].Next = &lists[2]
	lists[2].Next = &lists[3]
	lists[3].Next = &lists[4]
	lists[4].Next = &lists[5]

	for i := 0; i < len(lists); i++ {
		fmt.Printf("addr of list[%d] : %p\n", i, &lists[i])
	}

	p := partition(&lists[0], 3)
	printList(p)
}