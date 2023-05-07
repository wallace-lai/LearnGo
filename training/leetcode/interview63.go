// 链表

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms / 2.2MB
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	vecData := make([]*ListNode, 0)
// 	for p := head; p != nil; p = p.Next {
// 		vecData = append(vecData, p)
// 	}

// 	m := len(vecData) - n + 1
// 	if m == 1 {
// 		return vecData[0].Next
// 	}

// 	vecData[m - 2].Next = vecData[m - 1].Next
// 	return vecData[0]
// }

// v2 : 0ms / 2.1MB
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode {0, head}
	var (
		slow *ListNode = &dummy
		fast *ListNode = &dummy
	)

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

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

	p := removeNthFromEnd(&lists[0], 1)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}	
}