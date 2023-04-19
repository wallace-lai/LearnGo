package main

import "fmt"

// 方法2：使用const和iota定义枚举常量
const (
	BEIJING = iota	// iota = 0
	SHANGHAI		// iota = 1
	SHENZHEN		// iota = 2
)

const (
	a, b = iota + 1, iota + 2	// iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d						// iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f						// iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota * 3	// iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	i, k						// iota = 4, i = iota * 2, k = iota * 3, i = 8, k = 12
)
func main() {
	// 方法1：定义一个常量
	const PI float64 = 3.1415926
	fmt.Println("PI =", PI)
}
