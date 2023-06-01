// 堆

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/ipo/?envType=study-plan-v2&envId=top-interview-150

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

// v1 : 220ms
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)

	type Item struct {c, p int}
	array := make([]Item, n)
	for i := 0; i < n; i++ {
		array[i] = Item{capital[i], profits[i]}
	}
	sort.Slice(array, func(i, j int) bool {return array[i].c < array[j].c})

	h := &MaxHeap{}
	heap.Init(h)

	idx := 0
	for i := 0; i < k; i++ {
		// 先将启动资金小于w的项目压入堆中
		for idx < n && array[idx].c <= w {
			heap.Push(h, array[idx].p)
			idx++
		}
		if h.Len() == 0 {
			break
		}

		// 选取所有可选项目中获利最大的
		w += heap.Pop(h).(int)
	}

	return w
}

func main() {
	p := []int {
		1, 2, 3,
	}

	c := []int {
		0, 1, 1,
	}

	fmt.Println(findMaximizedCapital(3, 0, p, c))
}

