package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex	string `info:"sex"`
}

func FindTag(i interface {}) {
	t := reflect.TypeOf(i).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info :", tagInfo)
		fmt.Println("doc :", tagDoc)
	}
}

func main() {
	var r Resume

	FindTag(&r)
}