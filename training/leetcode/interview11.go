// 数组

package main

import (
	"math/rand"
)

// https://leetcode.cn/problems/insert-delete-getrandom-o1/?envType=study-plan-v2&id=top-interview-150

type RandomizedSet struct {
	vecData []int
	mapData map[int]int
}

// v1 : 160ms
func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}


func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mapData[val]; ok {
		return false
	}
	this.mapData[val] = len(this.vecData)
	this.vecData = append(this.vecData, val)
	return true
}


func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.mapData[val]
	if !ok {
		return false
	}

	lastIdx := len(this.vecData) - 1
	this.vecData[idx] = this.vecData[lastIdx]
	this.mapData[this.vecData[idx]] = idx
	this.vecData = this.vecData[:lastIdx]
	delete(this.mapData, val)
	return true
}


func (this *RandomizedSet) GetRandom() int {
	return this.vecData[rand.Intn(len(this.vecData))]
}

func main() {

}