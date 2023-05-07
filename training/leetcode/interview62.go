// 链表

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode.cn/problems/reverse-nodes-in-k-group/?envType=study-plan-v2&id=top-interview-150

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := ListNode {0, head}
	count := 0
	p := &dummy

	// 往前遍历k个节点
	for p.Next != nil && count < k {
		p = p.Next
		count++
	}

	// 不足k个节点的链表直接返回
	if count < k {
		return dummy.Next
	}

	// 反转后续的链表
	tmpNext := reverseKGroup(p.Next, k)

	// 反转当前的链表并将后续链表反转后的结果接到当前链表后面
	p.Next = nil
	dummy.Next = reverseList(dummy.Next, tmpNext)

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

	p := reverseKGroup(&lists[0], 2)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}	
}