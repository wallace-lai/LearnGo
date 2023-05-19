// 多维动态规划

package main

// https://leetcode.cn/problems/edit-distance/?envType=study-plan-v2&envId=top-interview-150

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// v1 : 8ms
// f[i][j]表示word1前i个字符转变成word2前j个字符所需的编辑距离
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	f := make([][]int, len1 + 1, len1 + 1)
	for i := 0; i < len1 + 1; i++ {
		f[i] = make([]int, len2 + 1, len2 + 1)
	}

	for i := 0; i <= len1; i++ {
		f[i][0] = i
	}
	for i := 0; i <= len2; i++ {
		f[0][i] = i
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i - 1] == word2[j - 1] {
				f[i][j] = f[i - 1][j - 1]
			} else {
				f[i][j] = getMin(f[i - 1][j - 1] + 1, getMin(f[i - 1][j] + 1, f[i][j - 1] + 1))
			}
		}
	}

	return f[len1][len2]
}
