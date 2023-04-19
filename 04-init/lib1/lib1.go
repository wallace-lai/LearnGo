package lib1

import (
	"fmt"
)

// 包提供的API
func Lib1Test() {
	fmt.Println("LIB1 TEST CALLED")
}

// 包的init方法
func init() {
	fmt.Println("LIB1 INIT CALLED")
}
