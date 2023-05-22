// 堆

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/kth-largest-element-in-an-array/?envType=study-plan-v2&envId=top-interview-150

// v1 :
// 112ms - O(NlogN)
// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	n := len(nums)
// 	return nums[n - k]
// }

// v2 : 76ms
// 76ms - O(N)
func quickSort(nums *[]int, left int, right int) {
	if left >= right {
		return
	}

	pivot := (*nums)[left]
	beg := left
	end := right

	for beg < end {
		for end > beg && (*nums)[end] >= pivot {
			end--
		}
		(*nums)[beg] = (*nums)[end]

		for beg < end && (*nums)[beg] <= pivot {
			beg++
		}
		(*nums)[end] = (*nums)[beg]
	}
	(*nums)[beg] = pivot

	quickSort(nums, left, beg - 1)
	quickSort(nums, beg + 1, right)
}

func randomPartition(nums *[]int, left int, right int) int {
	pivot := (*nums)[left]
	beg := left
	end := right

	for beg < end {
		for end > beg && (*nums)[end] >= pivot {
			end--
		}
		(*nums)[beg] = (*nums)[end]

		for beg < end && (*nums)[beg] <= pivot {
			beg++
		}
		(*nums)[end] = (*nums)[beg]
	}
	(*nums)[beg] = pivot

	return beg
}

func quickSelect(nums *[]int, left int, right int, k int) int {
	p := randomPartition(nums, left, right)
	if p == k {
		return (*nums)[p]
	} else if p < k {
		return quickSelect(nums, p + 1, right, k)
	} else {
		return quickSelect(nums, left, p - 1, k)
	}
}

// func findKthLargest(nums []int, k int) int {
// 	return quickSelect(&nums, 0, len(nums) - 1, len(nums) - k)
// }

// v3 :
// 80ms - O(NlogN) 
func swap(nums *[]int, beg int, end int) {
	tmp := (*nums)[beg - 1]
	(*nums)[beg - 1] = (*nums)[end - 1]
	(*nums)[end - 1] = tmp
}

func down(a *[]int, i int) {
	size := len(*a)

	left := i * 2
	right := i * 2 + 1
	largest := i

	if left <= size && (*a)[left - 1] > (*a)[largest - 1] {
		largest = left
	}
	if right <= size && (*a)[right - 1] > (*a)[largest - 1] {
		largest = right
	}

	if largest != i {
		swap(a, i, largest)
		down(a, largest)
	}
}

func up(a *[]int, i int) {
	for i / 2 > 0 && (*a)[i - 1] > (*a)[i / 2 - 1] {
		swap(a, i, i / 2)
		i = i / 2
	}
}

func buildHeap(a *[]int) {
	size := len(*a)
	for i := size / 2; i >= 1; i-- {
		down(a, i)
	}
}

func findKthLargest(nums []int, k int) int {
	buildHeap(&nums)
	for i := 0; i < k - 1; i++ {
		swap(&nums, 1, len(nums))
		nums = nums[0 : len(nums) - 1]
		down(&nums, 1)
	}

	return nums[0]
}

func main() {
	// nums := []int {
	// 	3, 2, 3, 1, 2, 4, 5, 5, 6,
	// }

	nums := []int {
		3, 2, 1, 5, 6, 4,
	}

	fmt.Println(findKthLargest(nums, 1))
	fmt.Println(findKthLargest(nums, 2))
	fmt.Println(findKthLargest(nums, 3))
	fmt.Println(findKthLargest(nums, 4))
	fmt.Println(findKthLargest(nums, 5))
	fmt.Println(findKthLargest(nums, 6))
}
