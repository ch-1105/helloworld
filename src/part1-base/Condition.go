package main

import "fmt"

func main() {
	// if 就是没括号的java
	if true {
		fmt.Println("true")
	} else if 1 > 0 {
		fmt.Println("1 > 0")
	} else {
		fmt.Println("false")
	}

	// switch
	switch 1 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	// for
	// 回顾下九九乘法表~
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, "*", j, "=", i*j, "\t")
		}
		if i < 9 {
			fmt.Println()
		}
	}
	fmt.Println()
	// for range
	for _, value := range "hello world" {
		fmt.Printf("%c ", value)
	}

	// goto 实现跳转
	// 多用于错误处理，但很少用
	fmt.Println("哈哈")
}
