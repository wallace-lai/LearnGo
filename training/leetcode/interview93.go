// 图

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/course-schedule/?envType=study-plan-v2&id=top-interview-150

// v1 : 8ms
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构造有向图
	indeg := make([]int, numCourses)
	graph := make([][]int, numCourses, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < len(prerequisites); i++ {
		from := prerequisites[i][1]
		to := prerequisites[i][0]
		graph[from] = append(graph[from], to)
		indeg[to]++
	}

	// 宽度优先搜索
	idx := 0
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)
	for len(queue) - idx > 0 {
		from := queue[idx]
		idx++

		result = append(result, from)
		for i := 0; i < len(graph[from]); i++ {
			to := graph[from][i]
			indeg[to]--
			if indeg[to] == 0 {
				queue = append(queue, to)
			}
		}
	}

	if len(result) != numCourses {
		return false
	}

	return true
}

func main() {

}