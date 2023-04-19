package main

import (
	"fmt"
)

func main () {
	var numbers = make([] int, 3, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 追加一个元素导致容量变大为原先的2倍
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	array := [] int { 1, 2, 3, 4, 5, 6, 7, 8 }
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(array), cap(array), array)

	// 打印切片，索引范围是[1, 4)
	fmt.Println(array[1:4])

	// 打印切片，索引范围是[0, 3)
	fmt.Println(array[:3])

	// 打印切片，索引范围是[3, len(array))
	fmt.Println(array[3:])

	// 打印切片，所有元素
	fmt.Println(array[:])

	// 使用copy进行深拷贝
	s2 := make([] int, 3)	// 给s2分配空间
	copy(s2, array)			// 将array中的数据拷贝到s2中
	fmt.Println(s2)
}