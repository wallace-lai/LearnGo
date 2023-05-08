// 链表

package main

import (
	"fmt"
)

// https://leetcode.cn/problems/lru-cache/?envType=study-plan-v2&id=top-interview-150

// v1 : 436ms
type Node struct {
	Key 	int
	Value 	int
	Prev 	*Node
	Next 	*Node
}

type LRUCache struct {
	MapData		map[int] *Node
	Head		*Node
	Tail		*Node
	Size		int
	Capacity	int
}

func (this *LRUCache) addToHead(node *Node) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) removeTail() *Node{
	node := this.Tail.Prev
	this.removeNode(node)
	return node
}

func Constructor(capacity int) LRUCache {
	inst := LRUCache {
		MapData: make(map[int]*Node),
		Head: &Node {0, 0, nil, nil},
		Tail: &Node {0, 0, nil, nil},
		Size: 0,
		Capacity: capacity,
	}
	inst.Head.Next = inst.Tail
	inst.Tail.Prev = inst.Head

	return inst
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.MapData[key]; !ok {
		return -1
	}

	node := this.MapData[key]
	this.moveToHead(node)
	return node.Value
}


func (this *LRUCache) Put(key int, value int)  {
	_, ok := this.MapData[key]
	if !ok {
		node := Node {Key: key, Value: value, Prev: nil, Next: nil}
		this.MapData[key] = &node
		this.addToHead(&node)
		this.Size++
		if this.Size > this.Capacity {
			removed := this.removeTail()
			delete(this.MapData, removed.Key)
			this.Size--
		}
	} else {
		node := this.MapData[key]
		node.Value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) dumpLRUCache() {
	fmt.Println("MapData =", this.MapData)
	fmt.Println("Size =", this.Size)
	fmt.Println("Capacity =", this.Capacity)
	fmt.Println("LinkedList :")
	for p := this.Head.Next; p != this.Tail; p = p.Next {
		fmt.Printf("(k = %d, v = %d)\n", p.Key, p.Value)
	}
}

func main() {
	inst := Constructor(2)
	inst.Put(1, 1)
	inst.Put(2, 2)
	// inst.dumpLRUCache()
	fmt.Println(inst.Get(1))
	// inst.dumpLRUCache()
	inst.Put(3, 3)
	fmt.Println(inst.Get(2))
	inst.Put(4, 4)
	fmt.Println(inst.Get(1))
	fmt.Println(inst.Get(3))
	fmt.Println(inst.Get(4))
}