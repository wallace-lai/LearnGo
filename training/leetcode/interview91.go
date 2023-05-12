// 图

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/clone-graph/?envType=study-plan-v2&id=top-interview-150

type Node struct {
	Val int
	Neighbors []*Node
}

// v1 : 0ms
func deepCopyGraph(node *Node, mapData *map[int]*Node) *Node {
	nodeVal := node.Val
	if v, ok := (*mapData)[nodeVal]; ok {
		return v
	}

	copyNode := &Node {Val: nodeVal, Neighbors: make([]*Node, 0)}
	(*mapData)[nodeVal] = copyNode

	for i := 0; i < len(node.Neighbors); i++ {
		copyNode.Neighbors = append(copyNode.Neighbors, deepCopyGraph(node.Neighbors[i], mapData))
	}

	return copyNode
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

    mapData := make(map[int]*Node)
	return deepCopyGraph(node, &mapData)
}

func dumpGraph(node *Node) {
	return
}

func main() {
	graph := []Node {
		{1, make([]*Node, 0)},
		{2, make([]*Node, 0)},
		{3, make([]*Node, 0)},
		{4, make([]*Node, 0)},
	}

	graph[0].Neighbors = []*Node {&graph[1], &graph[3]}
	graph[1].Neighbors = []*Node {&graph[0], &graph[2]}
	graph[2].Neighbors = []*Node {&graph[1], &graph[3]}
	graph[3].Neighbors = []*Node {&graph[0], &graph[2]}

	ret := cloneGraph(&graph[0])
	dumpGraph(ret)
}