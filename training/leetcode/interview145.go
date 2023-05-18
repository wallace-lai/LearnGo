// 多维动态规划

package main

// https://leetcode.cn/problems/longest-palindromic-substring/?envType=study-plan-v2&id=top-interview-150

// v1 : 56ms
// func longestPalindrome(s string) string {
// 	n := len(s)
// 	if n < 2 {
// 		return s
// 	}

// 	maxLen := 1
// 	begin := 0

// 	f := make([][]bool, n, n)
// 	for i := 0; i < n; i++ {
// 		f[i] = make([]bool, n, n)
// 	}

// 	for i := 0; i < n; i++ {
// 		f[i][i] = true
// 	}
// 	for l := 2; l <= n; l++ {
// 		for i := 0; i < n; i++ {
// 			j := l + i - 1
// 			if j >= n {
// 				break
// 			}

// 			if s[i] == s[j] {
// 				if l == 2 {
// 					f[i][j] = true
// 				} else {
// 					f[i][j] = f[i + 1][j - 1]
// 				}
// 			}

// 			if f[i][j] && l > maxLen {
// 				maxLen = l
// 				begin = i
// 			}
// 		}
// 	}

// 	return s[begin : begin + maxLen]
// }

// v2 : 4ms
func expand(s string, left, right int) (int, int) {
	n := len(s)
	for left >= 0 && right < n && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}

func longestPalindrome(s string) string {
	n := len(s)

	maxLen := 0
	begin := 0

	for i := 0; i < n; i++ {
		left, right := expand(s, i, i)
		if right - left + 1 > maxLen {
			maxLen = right - left + 1
			begin = left
		}

		left, right = expand(s, i, i + 1)
		if right - left + 1 > maxLen {
			maxLen = right - left + 1
			begin = left
		}
	}

	return s[begin : begin + maxLen]
}