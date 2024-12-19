package StructDemo

import "fmt"

// 结构体

type student struct {
	name string
	age  int
	// 结构体嵌套
	classTemp Class
	classmate
}

func (s student) say() {
	fmt.Println("学生输出")
}

type Class struct {
	ClassName string
	ClassNum  int
}

type classmate struct {
	classmateName string
	classmateAge  int
}

func main() {
	// type 关键字说明 1.定义结构体 2.定义接口 3.定义别名
	//	4.基于已有类型进行类型定义 5.类型判断

	// 初始化
	s1 := student{"zhangsan", 18,
		Class{
			ClassName: "一班",
			ClassNum:  20,
		},
		classmate{}}
	s2 := student{
		name: "lisi",
		age:  19,
		classTemp: Class{
			ClassName: "二班",
			ClassNum:  30,
		},
		classmate: classmate{"zhaoliu", 19},
	}
	fmt.Println(s1.classTemp.ClassNum) // 结构体使用非匿名时，需要多个点
	fmt.Println(s2.classmateName)      // 匿名属性直接加.访问

	fmt.Println(s1, s2)

	var stuList = []student{
		s1, s2,
	}
	fmt.Println(stuList)

	// 匿名结构体
	var stu2 = struct {
		className string
		classNum  int
	}{
		className: "一班",
		classNum:  20,
	}
	fmt.Println(stu2)

	// 结构体定义方法 在定义函数的基础上添加结构体绑定，且方法定义在结构体之外
	s1.say()

	// 通过方法传递结构体指针，避免结构体太大导致的拷贝以及在方法中修改这个结构体
}
