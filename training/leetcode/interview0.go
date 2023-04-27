package main

import (
	"fmt"
)

// https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&id=top-interview-150

// 4ms
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	pos := m + n - 1
// 	idx1 := m - 1
// 	idx2 := n - 1

// 	for idx1 >= 0 && idx2 >= 0 {
// 		if nums1[idx1] >= nums2[idx2] {
// 			nums1[pos] = nums1[idx1]
// 			idx1--
// 		} else {
// 			nums1[pos] = nums2[idx2]
// 			idx2--
// 		}
// 		pos--
// 	}

// 	for idx1 >= 0 {
// 		nums1[pos] = nums1[idx1]
// 		idx1--
// 		pos--
// 	}
// 	for idx2 >= 0 {
// 		nums1[pos] = nums2[idx2]
// 		idx2--
// 		pos--
// 	}
// }

// 0ms
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] >= nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}

	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

/*
（一）声明变量的几种方式
var a int		// 方法一：声明一个变量，默认值为0
var a int = 1	// 方法二：声明一个变量并初始化

var a = 1		// 方法三：声明变量时忽略类型，根据赋值类型自动推导
a := 1			// 方法四：省略var关键字，根据赋值类型自动推导【最常用】

				// 如何区别方法三和方法四？
				// 原则：有var关键字则无需冒号，否则需要冒号

var xx, yy = 100, "abcd"		// 方法五：一行代码同时声明多个变量
xx, yy := 100, "abcd"			// 根据上述原则，这个和方法五也是等价的

var (			// 方法六：同时声明多个变量的多行写法
	xx = 100
	yy = "abcd"
)

（二）声明切片的几种方式
array := int[] {1, 2, 3}		// 方法一：声明并初始化【常用】

var array []int					// 方法二：先声明切片，再使用make为其开辟空间，默认值为0
array = make([]int, 3)

var array []int = make([]int, 3)	// 方法三：将方法二放在一行

array := make([]int, 3)			// 方法四：声明切片并使用make为其开辟空间，使用推导的方式【常用】
array[0] = 1
array[1] = 2
array[2] = 3
*/

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)

	// nums1 := []int {1}
	// nums2 := []int {}
	// merge(nums1, 1, nums2, 0)

	// nums1 := []int {0}
	// nums2 := []int {1}
	// merge(nums1, 0, nums2, 1)

	fmt.Println(nums1)
}
