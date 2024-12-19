package main

import "fmt"

type duck interface {
	gaga()
}

type kedaya struct {
}

func (k *kedaya) gaga() {
	fmt.Println("gaga")
}

func mainInterface() {
	// go的接口具有多态效果
	var yaya duck = &kedaya{}
	yaya.gaga()

	/**
	使用：
		1. 声明接口
			type duck interface {
				gaga()
			}
		2. 创建结构体
			type kedaya struct {
			}
		3. 实现接口方法，并绑定结构体
			func (k *kedaya) gaga() {
				fmt.Println("gaga")
			}
		4. main中使用，声明接口变量，使用结构体对象
			var yaya duck = &kedaya{}
			yaya.gaga()
	*/
}
