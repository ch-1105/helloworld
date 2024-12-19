package main

import "fmt"

func main() {
	// 所以常用make进行初始化
	var a []int                // nil
	var slice = make([]int, 0) // empty
	if a == nil && slice != nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
