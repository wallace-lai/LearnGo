package main

import (
	"fmt"
)

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human Eat() ...")
}

func (this *Human) Walk() {
	fmt.Println("Human Walk() ...")
}

// 继承Human
type SuperMan struct {
	// 将基类名字Human作为类成员表示继承了Human
	Human
	level int
}

// 重写父类的方法Eat
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat() ...")
}

// 给子类新增方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly() ...")
}

func main() {
	h := Human{"zhangsan", "man"}
	h.Eat()
	h.Walk()

	// 定义并赋值子类对象，基类的赋值语法很怪
	// s := SuperMan{Human{"li4", "men"}, 88}
	var s SuperMan
	s.name = "li4"
	s.sex = "men"
	s.level = 88
	
	s.Walk() // 父类方法
	s.Eat()  // 子类重写后的方法
	s.Fly()  // 子类新增的方法
}
