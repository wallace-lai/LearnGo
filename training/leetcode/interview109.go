// 分治

package main

// https://leetcode.cn/problems/sort-list/?envType=study-plan-v2&id=top-interview-150

type ListNode struct {
	Val int
	Next *ListNode
}

// v1 : 48ms
// func merge(a *ListNode, b *ListNode) *ListNode {
// 	dummy := ListNode {0, nil}
// 	p := &dummy

// 	for a != nil && b != nil {
// 		if a.Val <= b.Val {
// 			p.Next = a
// 			a = a.Next
// 		} else {
// 			p.Next = b
// 			b = b.Next
// 		}
// 		p = p.Next
// 	}

// 	p.Next = a
// 	if b != nil {
// 		p.Next = b
// 	}

// 	return dummy.Next
// }

// func rangeMerge(lists []*ListNode, left int, right int) *ListNode {
// 	if left + 1 == right {
// 		return merge(lists[left], lists[right])
// 	}
// 	if left == right {
// 		return lists[left]
// 	}

// 	mid := left + (right - left) / 2
// 	return merge(rangeMerge(lists, left, mid), rangeMerge(lists, mid + 1, right))
// }

// func sortList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	lists := []*ListNode {}

// 	p := head
// 	for p != nil {
// 		tail := p.Next
// 		p.Next = nil
// 		lists = append(lists, p)
// 		p = tail
// 	}

// 	return rangeMerge(lists, 0, len(lists) - 1)
// }

// v2 : 52ms
func merge(a *ListNode, b *ListNode) *ListNode {
	dummy := ListNode {0, nil}
	p := &dummy

	for a != nil && b != nil {
		if a.Val <= b.Val {
			p.Next = a
			a = a.Next
		} else {
			p.Next = b
			b = b.Next
		}
		p = p.Next
	}

	p.Next = a
	if b != nil {
		p.Next = b
	}

	return dummy.Next
}

func rangeSort(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow

	return merge(rangeSort(head, mid), rangeSort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return rangeSort(head, nil)
}

func main() {
	nodes := []ListNode {
		{-1, nil},
		{5, nil},
		{3, nil},
		{4, nil},
		{0, nil},
	}
	nodes[0].Next = &nodes[1]
	nodes[1].Next = &nodes[2]
	nodes[2].Next = &nodes[3]
	nodes[3].Next = &nodes[4]

	sortList(sortList(&nodes[0]))
}