package main

import (
	"fmt"
	"time"
	// "time"
)

func main() {
	// 创建带缓存的channel
	c := make(chan int, 3)
	fmt.Println("len(c) =", len(c), ", cap =", cap(c))

	go func() {
		defer fmt.Println("Sub Go Exit...")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Printf("Sub Go is running, send item %d, len(c) = %d, cap(c) = %d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("Main Go Exit...")
}
