// 分治

package main

// https://leetcode.cn/problems/construct-quad-tree/?envType=study-plan-v2&id=top-interview-150

type Node struct {
	Val 		bool
	IsLeaf 		bool
	TopLeft 	*Node
	TopRight 	*Node
	BottomLeft 	*Node
	BottomRight *Node
}

// v1 : 8ms
func build(grid *[][]int, r0, c0, r1, c1 int) *Node {
	for row := r0; row < r1; row++ {
		for col := c0; col < c1; col++ {
			if (*grid)[row][col] != (*grid)[r0][c0] {
				topLeft := build(grid, r0, c0, (r0 + r1) / 2, (c0 + c1) / 2)
				topRight := build(grid, r0, (c0 + c1) / 2, (r0 + r1) / 2, c1)
				bottomLeft := build(grid, (r0 + r1) / 2, c0, r1, (c0 + c1) / 2)
				bottomRight := build(grid, (r0 + r1) / 2, (c0 + c1) / 2, r1, c1)
				return &Node {true, false, topLeft, topRight, bottomLeft, bottomRight}
			}
		}
	}

	return &Node {(*grid)[r0][c0] == 1, true, nil, nil, nil, nil}
}

func construct(grid [][]int) *Node {
	n := len(grid)
	return build(&grid, 0, 0, n, n)
}