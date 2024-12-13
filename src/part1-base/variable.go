package main

//func main() {
//	// go是静态语言，变量类型在编译期就确定，不能在运行期动态改变
//	// 变量首先需要定义，且需要使用类型，类型不能改变
//	// 定义变量的方式
//	// go定义了变量后必须使用~
//	// 1. var a int = 10
//	// 2. a := 10 海象赋值法
//
//	// 多变量定义
//	var usr1, usr2, usr3 string = "zhangsan", "lisi", "wangwu"
//	fmt.Println(usr1, usr2, usr3)
//
//	// go变量定义未赋值与Java一样 有零值
//
//	// 常量
//	const VARIABLE_A = 10
//	const VARIABLE_B int = 20
//	// iota 特殊常量 看起来就是个内存中计数器，在每次调用进行+1 中断后也会保存 初始值0  在另一个const块时 iota归0
//	const (
//		ERR1 = iota + 1
//		ERR2 = iota + 2
//		ERR3 = iota + 3
//		ERR4
//	)
//	fmt.Println(ERR1, ERR2, ERR3)
//
//	// 匿名变量  也就是当接受方法返回值有多个时，可能有些值我们并不需要，此时我们可以使用匿名变量_来接收
//	// 比如下列 只需要name
//	// 如 func returnTest()(int, string) {
//	// 	return 10, "hello world"
//	//	}
//	// _, name := returnTest(1,"")
//
//}
