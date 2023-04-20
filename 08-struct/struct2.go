package main

import (
	"fmt"
)

// 如果类名首字母是大写的表示其他的包也能访问这个类
// 同理如果成员首字母是大写的表示该成员是能被外界访问的
type Hero struct {
	Name string		// 对外可见
	Ad int			// 对外可见
	level int		// 对外不可见
}

// 首字母大写表示该方法对外可见
func (this Hero) GetName() string {
	return this.Name
}

// *表示this是指向对象的指针，无*表示this只是对象的一个拷贝
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

// 对于读操作可以不带*号但是对于写操作则必须带上否则修改不生效
func (this Hero) Show() {
	fmt.Println("Name =", this.Name)
	fmt.Println("Ad =", this.Ad)
	fmt.Println("Level =", this.level)
}

func main() {
	// 创建一个对象
	hero := Hero { Name: "zhang", Ad: 100, level: 1 }

	hero.Show()
	hero.SetName("LI")
	hero.Show()
}