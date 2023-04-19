package main

import (
	// lib1包前面加下划线表示匿名导包
	_ "LearnGo/04-init/lib1"

	// 可以给lib2包起别名mylib2
	// mylib2 "LearnGo/04-init/lib2"

	// 将lib2中的符号全部导入到当前的main包中
	// 类似C++中的using namespace std
	. "LearnGo/04-init/lib2"
)

func main() {
	// 匿名导包时可以不调用包中的方法但是包的init
	// 方法会被调用，即包可以正常以匿名的方式导入
	// lib1.Lib1Test()

	// 以别名的方式调用lib2中的方法
	// mylib2.Lib2Test()

	// 将lib2中符号全部导入当前main包中后可以直接调用包中方法
	Lib2Test()
}
