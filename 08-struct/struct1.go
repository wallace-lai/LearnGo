package main

import (
	"fmt"
)

// 使用type声明现有数据类型的别名
type myint int

// 使用type定义结构体类型
type Book struct {
	name string
	auth string
}

func printBook(book Book) {
	fmt.Println("Book Name : ", book.name)
	fmt.Println("Book Auth : ", book.auth)
}

// 注意结构体做函数入参是值传递
func changeBook(book Book) {
	book.auth = "666"
	book.name = "666"
}

// 传递指针才能起到修改作用
func realChangeBook(book * Book) {
	book.auth = "777"
	book.name = "777"
}

func main() {
	// var a myint = 10
	// fmt.Printf("type of a = %T, a = %d\n", a, a)

	var book1 Book
	book1.name = "Gone With The Wind"
	book1.auth = "Margaret Mitchell"

	fmt.Printf("type of book1 = %v\n", book1)
	printBook(book1)

	changeBook(book1)
	printBook(book1)

	realChangeBook(&book1)
	printBook(book1)
}
