// 区间

package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/?envType=study-plan-v2&id=top-interview-150

type pointSlice [][]int

func (p pointSlice) Len() int {
	return len(p)
}

func (p pointSlice) Less(lhs, rhs int) bool {
	return p[lhs][0] < p[rhs][0]
}

func (p pointSlice) Swap(lhs, rhs int) {
	tmp := p[lhs]
	p[lhs] = p[rhs]
	p[rhs] = tmp
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// v1 : 200ms
func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	p := pointSlice(points)
	sort.Sort(p)
	// fmt.Println("after sort p =", p)

	beg := p[0][0]
	end := p[0][1]
	result := 1
	for i := 1; i < len(p); i++ {
		// 和当前范围[beg, end]存在交集，更新交集
		if end >= p[i][0] {
			beg = getMax(beg, p[i][0])
			end = getMin(end, p[i][1])
			continue
		}
		// 否则需要再射一支箭并重置[beg, end]
		result++
		beg = p[i][0]
		end = p[i][1]

	}

	return result
}

func main() {
	points := [][]int {
		{10, 16}, {2, 8}, {1, 6}, {7, 12},
	}

	// points := [][]int {
	// 	{1, 2}, {3, 4}, {5, 6}, {7, 8},
	// }

	// points := [][]int {
	// 	{1, 2}, {2, 3}, {3, 4}, {4, 5},
	// }

	// points := [][]int {
	// 	{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8},
	// }

	fmt.Println(findMinArrowShots(points))
}