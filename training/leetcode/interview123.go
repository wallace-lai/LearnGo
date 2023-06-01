// 堆

package main

import (
	"fmt"
	"container/heap"
)

// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums?envType=study-plan-v2&envId=top-interview-150

var array1 []int
var array2 []int

type Item struct {i, j int}
type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	x1, y1 := h[i].i, h[i].j
	x2, y2 := h[j].i, h[j].j
	return array1[x1] + array2[y1] < array1[x2] + array2[y2]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

// v1 : 176ms
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	array1 = nums1
	array2 = nums2

	m := len(nums1)
	n := len(nums2)

	h := MinHeap{}
	heap.Init(&h)
	for i := 0; i < m; i++ {
		heap.Push(&h, Item{i, 0})
	}

	result := [][]int{}
	for h.Len() > 0 && len(result) < k {
		item := heap.Pop(&h).(Item)
		i, j := item.i, item.j
		result = append(result, []int{nums1[i], nums2[j]})
		if j + 1 < n {
			heap.Push(&h, Item{i, j + 1})
		}
	}

	return result
}

func main() {
	nums1 := []int {
		1, 7, 11,
	}

	nums2 := []int {
		2, 4, 6,
	}

	fmt.Println(kSmallestPairs(nums1, nums2, 3))
}