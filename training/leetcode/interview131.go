// 数学

package main

// https://leetcode.cn/problems/palindrome-number/?envType=study-plan-v2&envId=top-interview-150

// v1 : 4ms
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x % 10 == 0 && x != 0{
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted * 10 + x % 10
		x /= 10
	}

	return x == reverted || x == reverted / 10
}