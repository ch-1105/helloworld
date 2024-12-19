package StructDemo

import (
	"fmt"
)

//	func function_name( [parameter list] ) [return_types] {
//	  函数体
//	}
func add(b, c int, a string) (sum int, str string) {
	sum = b + c
	str = a
	//这里因为在函数体初始化定义了名称，也就是sum 和 str , return 会默认返回
	return // => 等同于 return sum,str
}

func sub() (str string) {
	str = "这是减法"
	return
}
func cal(choose string) func() {
	switch choose {
	case "sub":
		return func() {
			fmt.Println("sub 调用")
		}
	default:
		return nil
	}
}

func increment(i int) func() int {
	local := i
	return func() int {
		local += 1
		return local
	}
}

func main2() {
	a, b := add(1, 2, "hello world")
	fmt.Println(a, b)
	// 1. 函数本身可以当做变量进行传递,也可以进行返回, go可以返回多个值
	// 进行传递
	function := add
	c, d := function(2, 4, "Hello World !")
	fmt.Println(c, d)
	// 进行返回
	cal("sub")() //注意双括号，以函数方式进行调用
	// 2. 匿名函数 闭包
	// 感觉是在不同作用块中，每个变量独立副本
	/**
	这部分代码每次调用只会i+1  无论循环多少次都是
	func increment(i int) int {
		local := i
		return func () int {
			local += 1
			return local
		}() //最后加这个括号代表调用，不加的话只是定义函数
	}

	修改为函数调用后 for循环正常 下次想使用重新定义函数变量即可归0
	func increment(i int) func() int {
		local := i
		return func() int {
			local += 1
			return local
		}
	}
	*/
	incr := increment(0)
	for i := 0; i < 5; i++ {
		fmt.Println(incr())
	}
	// 3. 函数可以是接口

	// 4. ...any 接收任意长度参数，函数体等同于接收一个切片类型
}
