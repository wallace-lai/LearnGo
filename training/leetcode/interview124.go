// 堆

package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.cn/problems/find-median-from-data-stream/?envType=study-plan-v2&envId=top-interview-150

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Peek() any {
	return h[0]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

// v1 : 284ms
type MedianFinder struct {
	smallest	MinHeap
	largest		MinHeap
}

func Constructor() MedianFinder {
	f := MedianFinder {}
	heap.Init(&f.smallest)
	heap.Init(&f.largest)
	return f
}

func (this *MedianFinder) AddNum(num int)  {
	if this.smallest.Len() == 0 || num <= -this.smallest.Peek().(int) {
		heap.Push(&this.smallest, -num)
		if this.smallest.Len() > this.largest.Len() + 1 {
			top := heap.Pop(&this.smallest).(int)
			heap.Push(&this.largest, -top)
		}
	} else {
		heap.Push(&this.largest, num)
		if this.largest.Len() > this.smallest.Len() {
			top := heap.Pop(&this.largest).(int)
			heap.Push(&this.smallest, -top)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.smallest.Len() > this.largest.Len() {
		return float64(-this.smallest.Peek().(int))
	}
	
	return (float64(this.largest.Peek().(int)) + float64(-this.smallest.Peek().(int))) / 2.0
}

func main() {
	nums := []int {
		// 6, 10, 2, 6, 5, 0,
		1, 2, 3,
	}

	f := Constructor()

	length := len(nums)
	for i := 0; i < length; i++ {
		f.AddNum(nums[i])
		fmt.Println("i =", i)
		fmt.Println("smallest =", f.smallest)
		fmt.Println("largest =", f.largest)
		fmt.Println()
	}
	fmt.Println(f.FindMedian())
}