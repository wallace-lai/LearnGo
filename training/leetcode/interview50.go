// 区间

package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/insert-interval/?envType=study-plan-v2&id=top-interview-150

type intervalSlice [][]int

func (i intervalSlice) Len() int {
	return len(i)
}

func (i intervalSlice) Less(lhs, rhs int) bool {
	return i[lhs][0] < i[rhs][0]
}

func (i intervalSlice) Swap(lhs, rhs int) {
	tmp := i[lhs]
	i[lhs] = i[rhs]
	i[rhs] = tmp
}

func getMax(lhs, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	i := intervalSlice(intervals)
	sort.Sort(i)

	beg := i[0][0]
	end := i[0][1]
	result := make([][]int, 0, len(intervals))
	for idx := 1; idx < len(i); idx++ {
		if end >= i[idx][1] {
			// [i[idx][0], i[idx][1]] is covered by [beg, end], do nothing
		} else if end >= i[idx][0] {
			end = getMax(end, i[idx][1])
		} else {
			result = append(result, []int {beg, end})
			beg = i[idx][0]
			end = i[idx][1]
		}
	}
	result = append(result, []int {beg, end})

	return result
}

// v1 : 4ms
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

func main() {

}
