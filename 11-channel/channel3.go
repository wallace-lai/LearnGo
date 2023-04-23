package main

import (
	"fmt"
)

func main() {
	// 创建不带缓存的channel
	c := make(chan int)

	// 创建子Go程
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// 使用close关闭channel，如果这里不关闭channel则会导致主Go程死锁
		close(c)
	}()

	for {
		// ok为true表示channel没有关闭，否则表示channel已经关闭了
		data, ok := <- c
		if (ok) {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Goroutine Exit...")
}