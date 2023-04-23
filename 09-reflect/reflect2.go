package main

import (
	"fmt"
)

// 定义两个接口
type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

// 定义具体类型并实现两个接口
type Book struct {

}

// 因为Book里同时实现了Reader和Writer接口
// 所以Book它即是Reader又是Writer
func (this *Book) Read() {
	fmt.Println("Read a book")
}

func (this *Book) Write() {
	fmt.Println("Write a book")
}

func main() {
	// b : pair<type : Book, value : Book{}的地址>
	b := &Book{}

	// r : pair<type : , value : >
	var r Reader
	r = b		// r : pair<type : Book, value : Book{}的地址>
	r.Read()	// 调用的是Book类中的Read方法

	var w Writer
	w = r.(Writer)		// 此处断言为什么会成功，因为w和r具体的type是一致的都指向Book类型
	w.Write()

}