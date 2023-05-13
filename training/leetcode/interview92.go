// 图

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/evaluate-division/?envType=study-plan-v2&id=top-interview-150

// v1 : 0ms
func depthFirstSearch(graph *[][]float64, beg, end int, vis *[]byte, prod float64, result *float64, found *bool) {
	if beg == end {
		*found = true
		*result = prod
		return
	}

	(*vis)[beg] = 1
	size := len(*vis)
	for i := 0; i < size; i++ {
		if (*graph)[beg][i] > 0 && (*vis)[i] != 1 {
			depthFirstSearch(graph, i, end, vis, prod * (*graph)[beg][i], result, found)
		}
	}
	(*vis)[beg] = 0
}

func queryResult(graph *[][]float64, mapData *map[string]int, from string, to string) float64 {
	// 先检查待查询节点是否存在于图中
	if _, ok := (*mapData)[from]; !ok {
		return -1.0
	}
	if _, ok := (*mapData)[to]; !ok {
		return -1.0
	}

	// 深度优先搜索答案 
	result := 1.0
	beg := (*mapData)[from]
	end := (*mapData)[to]
	size := len(*mapData)
	vis := make([]byte, size, size)
	for i := 0; i < size; i++ {
		vis[i] = 0
	}

	found := false
	depthFirstSearch(graph, beg, end, &vis, 1.0, &result, &found)

	if !found {
		return -1.0
	}

	return result
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 先构建字符串到数组下标的映射
	idx := 0
	mapData := make(map[string]int)
	for i := 0; i < len(equations); i++ {
		if _, ok := mapData[equations[i][0]]; !ok {
			mapData[equations[i][0]] = idx
			idx++
		}
		if _, ok := mapData[equations[i][1]]; !ok {
			mapData[equations[i][1]] = idx
			idx++
		}
	}

	// 构建有向图
	graph := make([][]float64, len(mapData), len(mapData))
	for i := 0; i < len(mapData); i++ {
		graph[i] = make([]float64, len(mapData))
		for j := 0; j < len(mapData); j++ {
			graph[i][j] = -1.0
		}
	}

	for i := 0; i < len(equations); i++ {
		from := mapData[equations[i][0]]
		to := mapData[equations[i][1]]
		graph[from][from] = 1.0
		graph[to][to] = 1.0
		graph[from][to] = values[i]
		graph[to][from] = 1.0 / values[i]
	}

	// 搜索答案
	result := make([]float64, 0)
	for i := 0; i < len(queries); i++ {
		from := queries[i][0]
		to := queries[i][1]
		ret := queryResult(&graph, &mapData, from, to)
		result = append(result, ret)
	}

	return result
}

