// 链表

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// v1 : 8ms
// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}

// 	mapData := make(map[*ListNode]int)
// 	for head.Next != nil {
// 		if _, ok := mapData[head.Next]; ok {
// 			return true
// 		} else {
// 			mapData[head.Next] = 1
// 		}

// 		head = head.Next
// 	}

// 	return false
// }

// v2 : 4ms
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func main() {
	list := []ListNode {
		{1, nil}, {2, nil}, {3, nil}, {4, nil}, {5, nil},
	}

	list[0].Next = &list[1]
	list[1].Next = &list[2]
	list[2].Next = &list[3]
	list[3].Next = &list[4]
	// list[4].Next = nil
	list[4].Next = &list[0]

	fmt.Println(hasCycle(&list[0]))
}