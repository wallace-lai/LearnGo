package main

import (
	"fmt"
)

func main() {
	// pair<type : string, value : "abcd">
	a := "abcd"

	var genericType interface {}
	genericType = a		// pair<type : string, value : "abcd">

	value, _ := genericType.(string)
	fmt.Println(value)
}
