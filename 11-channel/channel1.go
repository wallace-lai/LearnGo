package main

import (
	"fmt"
)

// 创建channel
// make(chan Type)		// 等价于make(chan Type, 0)
// make(chan Type, capacity)

// 读写channel
// channel <- value	// 发送value到channel中
// <- channel			// 接受经由channel发送过来的数据并丢弃之
// x := <- channel		// 接收channel中的数据并存入x中
// x, ok := <- channel	// 功能同上，同时检查通道是否已关闭或者是否为空

func main() {
	// 定义一个无缓存的channel
	// var c chan int	// nil channel will cause deadlock
	c := make(chan int)

	go func() {
		defer fmt.Println("Goroutine exit...")

		fmt.Println("Goroutine running...")

		c <- 666	// 将666送入channel中
	}()

	num := <- c		// 接收channel中的数据
	fmt.Println("num =", num)
	fmt.Println("Main Goroutine exit...")
}
