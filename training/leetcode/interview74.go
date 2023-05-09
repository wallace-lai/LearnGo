// 二叉树

package main

type Node struct {
	Val 	int
	Left 	*Node
	Right 	*Node
	Next 	*Node
}

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/

// v1 : 0ms
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	idx := 0
	queue := make([]*Node, 0)

	root.Next = nil
	queue = append(queue, root)

	for len(queue) - idx > 0 {
		size := len(queue) - idx

		// 先扩展下一层元素
		for i := 0; i < size; i++ {
			curr := queue[idx + i]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		// 再给当前层元素设置右侧节点
		prev := queue[idx]
		for i := 1; i < size; i++ {
			curr := queue[idx + i]
			prev.Next = curr
			prev = curr
		}
		prev.Next = nil

		// 再将当前层元素出队列
		idx += size
	}

	return root
}

func main() {

}