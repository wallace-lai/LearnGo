package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/valid-palindrome/?envType=study-plan-v2&id=top-interview-150

func isAlnum(char byte) bool {
	return (char >= 'A' && char <= 'Z') ||
		(char >= 'a' && char <= 'z') ||
		(char >= '0' && char <= '9')
}

// 196ms
// func isPalindrome(s string) bool {
// 	// 过滤字母和数字之外的字符
// 	var filterStr string
// 	for i := 0; i < len(s); i++ {
// 		if isAlnum(s[i]) {
// 			filterStr += string(s[i])
// 		}
// 	}

// 	// 统一转换成大写字母
// 	filterStr = strings.ToUpper(filterStr)
// 	strLen := len(filterStr)
// 	if strLen == 0 {
// 		return true
// 	}

// 	// 双指针对比是否为回文串
// 	i := 0
// 	j := strLen - 1
// 	for i < j {
// 		if filterStr[i] != filterStr[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}

// 	return true
// }

// 0ms
func isPalindrome(s string) bool {
	upperStr := strings.ToUpper(s)
	strLen := len(upperStr)
	if strLen == 0 {
		return true
	}

	i := 0
	j := strLen - 1
	for i < j {
		if isAlnum(upperStr[i]) && isAlnum(upperStr[j]) {
			if upperStr[i] != upperStr[j] {
				return false
			}
			i++
			j--
		} else if isAlnum(upperStr[i]) && !isAlnum(upperStr[j]) {
			j--
		} else if !isAlnum(upperStr[i]) && isAlnum(upperStr[j]) {
			i++
		} else {
			i++
			j--
		}
	}

	return true
}

/*
一、标准库strings的常用API
（1）使用len(s)获取字符串的长度，仅限于ASCII编码
	str := "hello"		// len(str) = 5

	如果字符串里包含了汉字，即使用UTF-8编码的话，则len的结果不能表示真实的字符个数
	str := "hello世界"	// len(str) = 11，因为UTF-8中一个汉字需要3个字节

	如果想要正确得到字符串中的字符个数，需要使用rune数组，如下所示
	str := "hello世界"
	fmt.Println(len([]rune(str)))	// 结果为7

（2）可以使用[]去引用字符串中的字符，但是不可以修改
	str := "hello"
	isAlnum(str[0])		// 合法
	str[0] = 'H'		// 非法

（3）大小写转换
	func ToLower(s string) string
	func ToUpper(s string) string
*/

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))
}
