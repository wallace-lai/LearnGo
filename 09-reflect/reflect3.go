package main

import (
	"fmt"
	"reflect"
)

func reflectVariable(i interface {}) {
	fmt.Println("type :", reflect.TypeOf(i))
	fmt.Println("value :", reflect.ValueOf(i))
}

type User struct {
	Id int
	Name string
	Age int
}

func (this *User) Call() {
	fmt.Println("User is called...")
	fmt.Printf("%v\n", this)
}

func reflectDetail(input interface {}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType =", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue =", inputValue)

	// 根据input的type信息获取结构体里的字段
	// 1. 获取NumField
	// 2. 得到每个field
	// 3. 通过field有一个interface方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
	}
	
	fmt.Println("+++++++++++++++++++")
	
	// 根据input的type信息获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)
	}
}

func main() {
	// 获取基本类型变量的反射信息
	num := 3.1415926
	reflectVariable(num)

	fmt.Println("+++++++++++++++++++")
	// 获取自定义类型变量的反射信息
	user := User {1, "Wallace", 18}
	reflectDetail(user)
}