package main

import (
	"fmt"
)

func getMin(array []int) int {
	size := len(array)
	if size < 1 {
		return 0
	}

	min := array[0]
	for i := 1; i < size; i++ {
		if array[i] < min {
			min = array[i]
		}
	}

	return min
}

func getMax(array []int) int {
	size := len(array)
	if size < 1 {
		return 0
	}

	max := array[0]
	for i := 1; i < size; i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	return max
}


func equalFrequency(word string) bool {

}

func main() {
	fmt.Println(equalFrequency("abc"))
}