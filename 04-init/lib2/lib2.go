package lib2

import (
	"fmt"
)

// 包提供的API
func Lib2Test() {
	fmt.Println("LIB2 TEST CALLED")
}

// 包的init方法
func init() {
	fmt.Println("LIB2 INIT CALLED")
}
