// 链表

package main

import (
	"fmt"
)

type Node struct {
	Val int
	Next *Node
	Random *Node
}

// https://leetcode.cn/problems/copy-list-with-random-pointer/?envType=study-plan-v2&id=top-interview-150

// v1 : wrong answer
// func copyRandomList(head *Node) *Node {
// 	mapData := make(map[*Node]*Node)
// 	var vecData []Node

// 	p := head
// 	idx := 0
// 	for p != nil {
// 		vecData = append(vecData, Node {p.Val, nil, nil})
// 		mapData[p] = &vecData[idx]

// 		p = p.Next
// 		idx++
// 	}

// 	p = head
// 	idx = 0
// 	for p != nil {
// 		if p.Next != nil {
// 			vecData[idx].Next = mapData[p.Next]
// 		}
// 		if p.Random != nil {
// 			vecData[idx].Next = mapData[p.Next]
// 		}

// 		p = p.Next
// 		idx++
// 	}

// 	return &vecData[0]
// }

// v2 : 4ms
var mapData map[*Node]*Node

func copyRandomList(head *Node) *Node {
    if (head == nil) {
		return nil
	}

	if _, ok := mapData[head]; ok {
		return mapData[head]
	}

	node := &Node{head.Val, nil, nil}
	mapData[head] = node
	node.Next = copyRandomList(head.Next)
	node.Random = copyRandomList(head.Random)

	return node
}

func main() {

}

