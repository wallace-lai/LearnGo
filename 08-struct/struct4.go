package main

import (
	"fmt"
)

// 父类接口本质上是一个指针
type AnimalIntf interface {
	// 会跑会叫会吃东西的我们都可以理解为动物
	Run()  // 跑
	Eat()  // 进食
	Bark() // 叫
}

// 实现了所有父类接口的子类Cat
type Cat struct {
}

func (this *Cat) Run() {
	fmt.Println("Cat Run() ...")
}

func (this *Cat) Eat() {
	fmt.Println("Cat Eat() ...")
}

func (this *Cat) Bark() {
	fmt.Println("Cat Bark() ...")
}

// 实现了所有父类接口的子类Dog
type Dog struct {
}

func (this *Dog) Run() {
	fmt.Println("Dog Run() ...")
}

func (this *Dog) Eat() {
	fmt.Println("Dog Eat() ...")
}

func (this *Dog) Bark() {
	fmt.Println("Dog Bark() ...")
}

func main() {
	var animal AnimalIntf // 定义父类指针
	cat := Cat{}          // 子类对象
	dog := Dog{}          // 子类对象

	// 根据父类指针的指向调用子类对应的Bark方法
	animal = &cat
	animal.Bark()

	animal = &dog
	animal.Bark()
}
