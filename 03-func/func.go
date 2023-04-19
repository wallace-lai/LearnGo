package main

import (
	"fmt"
)

// 单个返回值函数
func foo1(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 100
}

// 多个匿名返回值函数
func foo2(a string, b int) (int, int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 100, 200
}

// 多个带形参名的返回值的函数
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("--- foo3 ---")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 形参在未赋值之前值为0
	fmt.Println("r1 = ", r1, " r2 = ", r2)

	// 注意带形参返回值的写法
	r1 = 100
	r2 = 200
	return
}

func main() {
	c := foo1("abcd", 100)
	fmt.Println("c =", c)

	d, f := foo2("hello world", 100)
	fmt.Println("d =", d, " f =", f)

	i, j := foo3("hello", 100)
	fmt.Println("i = ", i, " j = ", j)
}
