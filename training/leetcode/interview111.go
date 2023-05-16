// 分治

package main

// https://leetcode.cn/problems/merge-k-sorted-lists/?envType=study-plan-v2&id=top-interview-150

type ListNode struct {
	Val int
	Next *ListNode
}

// v1 : 8ms
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

func rangeMerge(lists []*ListNode, left int, right int) *ListNode {
	if left + 1 == right {
		return merge(lists[left], lists[right])
	}
	if left == right {
		return lists[left]
	}

	mid := left + (right - left) / 2
	return merge(rangeMerge(lists, left, mid), rangeMerge(lists, mid + 1, right))
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return rangeMerge(lists, 0, len(lists) - 1)
}