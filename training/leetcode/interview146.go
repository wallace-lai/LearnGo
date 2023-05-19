// 多维动态规划

package main

// https://leetcode.cn/problems/interleaving-string/?envType=study-plan-v2&envId=top-interview-150

// v1 : 0ms
// func isInterleave(s1 string, s2 string, s3 string) bool {
// 	len1 := len(s1)
// 	len2 := len(s2)
// 	len3 := len(s3)
// 	if len1 + len2 != len3 {
// 		return false
// 	}

// 	f := make([][]bool, len1 + 1, len1 + 1)
// 	for i := 0; i < len1 + 1; i++ {
// 		f[i] = make([]bool, len2 + 1, len2 + 1)
// 	}

// 	f[0][0] = true
// 	for i := 0; i <= len1; i++ {
// 		for j := 0; j <= len2; j++ {
// 			p := i + j - 1
// 			if i > 0 {
// 				f[i][j] = f[i][j] || (f[i - 1][j] && s1[i - 1] == s3[p])
// 			}
// 			if j > 0 {
// 				f[i][j] = f[i][j] || (f[i][j - 1] && s2[j - 1] == s3[p])
// 			}
// 		}
// 	}

// 	return f[len1][len2]
// }

// v2 : 0ms - 没搞懂
func isInterleave(s1 string, s2 string, s3 string) bool {
	len1 := len(s1)
	len2 := len(s2)
	len3 := len(s3)
	if len1 + len2 != len3 {
		return false
	}

	f := make([]bool, len2 + 1, len2 + 1)


	f[0] = true
	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			p := i + j - 1
			if i > 0 {
				f[j] = (f[j] && s1[i - 1] == s3[p])
			}
			if j > 0 {
				f[j] = f[j] || (f[j - 1] && s2[j - 1] == s3[p])
			}
		}
	}

	return f[len2]
}
