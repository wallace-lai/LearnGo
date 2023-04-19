package main

import (
	"fmt"
)

func main() {
	// 声明map1是string到string的映射
	var map1 map[string] string
	if (map1 == nil) {
		fmt.Println("map1 is null for now")
	}

	// 在使用map前需要使用make分配空间
	map1 = make(map[string] string, 10)

	map1["one"] = "1"
	map1["two"] = "2"

	fmt.Println(map1)

	// 使用make声明map但未分配空间
	map2 := make(map [int] string)
	map2[1] = "one"
	map2[2] = "two"
	fmt.Printf("len = %d, map = %v\n", len(map2), map2)

	map2[2] = "TWO"
	map2[3] = "three"
	fmt.Printf("len = %d, map = %v\n", len(map2), map2)

	map3 := map [int] string {
		1 : "one",
		2 : "two",
		6 : "six",
	}
	fmt.Println(map3)
}