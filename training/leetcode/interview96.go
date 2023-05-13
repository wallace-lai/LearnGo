// 图的广度优先搜索

package main

import "fmt"

// https://leetcode.cn/problems/minimum-genetic-mutation/?envType=study-plan-v2&id=top-interview-150

// v1 : timeout
// func minMutation(startGene string, endGene string, bank []string) int {
// 	// 构造基因哈希表
// 	mapData := make(map[string]bool)
// 	for _, gen := range bank {
// 		mapData[gen] = true
// 	}
// 	if _, ok := mapData[endGene]; !ok {
// 		return -1
// 	}

// 	idx := 0
// 	queue := make([]string, 0)
// 	vis := make(map[string]bool)

// 	result := -1
// 	queue = append(queue, startGene)
// 	vis[startGene] = true

// 	for len(queue) - idx > 0 {
// 		result++

// 		size := len(queue) - idx
// 		for i := 0; i < size; i++ {
// 			curr := queue[idx]
// 			idx++
// 			if curr == endGene {
// 				return result
// 			}

// 			for j := 0; j < len(curr); j++ {
// 				for _, gen := range "AGCT" {
// 					if curr[j] == byte(gen) {
// 						continue
// 					}
					
// 					next := curr
// 					if j == 0 {
// 						next = string(gen) + curr[j + 1:]
// 					} else if j == len(curr) - 1 {
// 						next = curr[0:len(curr)-1] + string(gen)
// 					} else {
// 						next = curr[0:j] + string(gen) + curr[j + 1:]
// 					}
// 					if _, ok := mapData[next]; ok {
// 						queue = append(queue, next)
// 						vis[next] = true
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return -1
// }

// v2 : 0ms
func checkValid(curr string, next string, vis *map[string]bool) bool {
	sum := 0
	for i := 0; i < len(curr); i++ {
		if curr[i] != next[i] {
			sum++
		}
	}
	if sum != 1 {
		return false
	}

	if _, ok := (*vis)[next]; ok {
		return false
	}

	return true
}

func minMutation(startGene string, endGene string, bank []string) int {
	idx := 0
	queue := make([]string, 0)
	vis := make(map[string]bool)

	result := -1
	queue = append(queue, startGene)
	vis[startGene] = true

	for len(queue) - idx > 0 {
		result++

		size := len(queue) - idx
		for i := 0; i < size; i++ {
			curr := queue[idx]
			idx++
			if curr == endGene {
				return result
			}

			for _, next := range bank {
				if checkValid(curr, next, &vis) {
					queue = append(queue, next)
					vis[next] = true
				}
			}
		}
	}

	return -1
}

func main() {
	startGen := "AAAAAAAA"
	endGen := "CCCCCCCC"
	bank := []string {
		"AAAAAAAA",
		"AAAAAAAC",
		"AAAAAACC",
		"AAAAACCC",
		"AAAACCCC",
		"AACACCCC",
		"ACCACCCC",
		"ACCCCCCC",
		"CCCCCCCA",
	}

	fmt.Println(minMutation(startGen, endGen, bank))
}