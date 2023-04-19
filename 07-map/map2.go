package main

import (
	"fmt"
)

func printMap(city map [string] string) {
	// 注意形参city是实参的一个引用
	for key, value := range city {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
}

func main() {
	city := make(map [string] string)

	// 插入
	city["China"] = "BeiJing"
	city["Japan"] = "Tokyo"
	city["USA"] = "NewYork"

	// 遍历
	for key, value := range city {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}

	// 删除
	delete(city, "USA")

	// 修改
	city["China"] = "ShangHai"

	printMap(city)
}