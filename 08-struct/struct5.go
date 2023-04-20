package main

import (
	"fmt"
)

// interface{}表示万能数据类型
func MyPrint(arg interface{} ) {
	fmt.Println("MyPrint called ...")
	fmt.Println(arg)

	// 使用interface提供的断言机制判断arg的具体类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value =", value)
		fmt.Printf("arg type is %T\n", value)
	}
	fmt.Println("=================================\n\n")
}

// 自定义类型
type Book struct {
	name string
	auth string
}

func main() {
	book := Book {"Hello", "666"}
	MyPrint(book)	// 处理自定义类型
	MyPrint(100)	// 处理整型数据
	MyPrint("abcd")	// 处理字符串类型数据
}

