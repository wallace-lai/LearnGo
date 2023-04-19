package main

import (
	"fmt"
)

func main() {
	// 声明slice1是一个切片且初始化，默认值为1,2,3则长度len为3
	slice1 := []int{1, 2, 3}

	// 声明slice2是一个切片但是没有给slice分配空间
	var slice2 []int
	// 使用make为slice开辟3个元素的空间
	slice2 = make([]int, 3)

	// 声明slice3是一个切片且给slice分配空间，默认值为0
	var slice3 []int = make([]int, 3)

	// 声明slice4是一个切片且给slice分配空间，但使用推导方式
	slice4 := make([]int, 4)

	// 判断slice是否已分配空间
	if slice2 == nil {
		fmt.Println("slice是一个空切片")
	} else {
		fmt.Println("slice是有空间的")
	}
}
