package main

import (
	"fmt"
	// "runtime"
	"time"
)

// func main() {
// 	// 用Go程跑无参匿名函数
// 	go func() {
// 		defer fmt.Println("A.defer")

// 		func() {
// 			defer fmt.Println("B.defer")

// 			// 退出当前的Go程
// 			runtime.Goexit()
// 			fmt.Println("B")
// 		}()

// 		fmt.Println("A")
// 	}()

// 	for {
// 		time.Sleep(1 * time.Second)
// 	}
// }

func main() {
	// 用Go程跑有参匿名函数，如何得到匿名函数的返回值?
	go func(a int, b int) bool {
		fmt.Printf("a = %d, b = %d\n", a, b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}