package main

import (
	"fmt"
)

func main() {
	// 方法1：声明一个变量，默认值为0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法2：声明一个变量并初始化
	var b int = 1
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)

	// 方法3：变量初始化时省略数据类型，通过值的自动匹配变量的类型
	var c = 100
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)

	var d = "abcd"
	fmt.Println("d =", d)
	fmt.Printf("type of d = %T\n", d)

	// 方法4：省略var关键字，直接自动匹配。只能用于函数内部局部变量的声明
	e := 200
	fmt.Println("e =", e)
	fmt.Printf("type of e = %T\n", e)

	// 同时声明多个变量的单行写法
	var xx, yy = 100, "abcd"
	fmt.Printf("type of xx = %T\n", xx)
	fmt.Printf("type of yy = %T\n", yy)

	// 同时声明多个变量的多行写法
	var (
		ii int  = 100
		jj bool = true
	)
	fmt.Printf("type of ii = %T\n", ii)
	fmt.Printf("type of jj = %T\n", jj)
}
