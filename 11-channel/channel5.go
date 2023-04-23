package main

import (
	"fmt"
)

// select的用法
// select {
// case <- chan1:
// 	// 如果chan1成功读到数据，则进行该case语句的处理
// case chan2 <- 1:
// 	// 如果成功向chan2写入数据，则进行该case语句的处理
// default:
// 	// default处理流程
// }

func Fib(c chan int, quit chan int) {
	x, y := 1, 1

	for {
		select {
		// 若c可写则更新x和y的值
		case c <- x:
			x = y
			y = x + y
		// 若quit可读则说明循环结束
		case <- quit:
			fmt.Println("Quit...")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<- c)
		}

		quit <- 0
	}()

	// channel做入参其实是传递的引用
	Fib(c, quit)
}
