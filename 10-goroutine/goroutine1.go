package main

import (
	"fmt"
	"time"
)

func NewTask() {
	i := 0
	for {
		i++
		fmt.Printf("New Go Routine i = %d\n", i)
		time.Sleep((1 * time.Second))
	}
}

func main() {
	// 创建一个goroutine运行NewTask
	go NewTask()

	i := 0
	for {
		i++
		fmt.Printf("Main Go Routine i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

