package container

import (
	"fmt"
)

func main3() {
	// slice  => java List
	// 使用append ， 接受得是它 输入参数也得是
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice)

	// 切片创建方法 1.从数组 2.string{} 3.make
	fromList := []int{1, 2, 3, 4, 5} //区别于数组就是 slice没有[num]
	fmt.Println(fromList)

	a := [5]string{"1", "2"}
	b := a[1:3] // 左闭右开 类似python
	fmt.Println(b)

	sliceMake := make([]int, 5)
	sliceMake[2] = 3
	fmt.Println(sliceMake) // 默认为0

	// 切片元素访问 访问多个元素【start : end】 左闭右开
	fmt.Println(sliceMake[1:3])

	// slice功能 1.copy 2.append

	// 底层原理
	// 在传递时是值传递还是引用传递
	// 实际是值传递，但是效果是引用传递
	// 原因就是 slice底层类型在传递时 传递了固定长度的复制后的数组 那么问题来了，为什么会把java传递回来
	// 是因为传递的指针也同时指向了同一个地址，但是在扩容之后两个指向的对象不相等了
	// 扩容机制：在数组小于1024 是 2倍
	//
	//strSlice := []string{"1", "2"}
	//printSlice(strSlice)
	//fmt.Println(strSlice)
	//func printSlice(data []string) {
	//	data[0] = "java" 0位可以修改为java，可以打印
	//	for i := 0; i < len(data); i++ {
	//		data = append(data, data[i]) 扩容之后，原方法打印数组不能打印
	//	}
	//}

	var num1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s1 := num1[2:5]
	s2 := num1[3:6]
	s2[1] = 8888
	fmt.Println(s1, s2)

	// 扩容机制
	// 如果新的切片长度小于1024，容量通常会翻倍。
	// 如果新的切片长度大于等于1024，容量的增长将会是 "旧容量 + 旧容量/4" （即增加25%）。
	for i := 0; i < 2048; i++ {
		num1 = append(num1, i)
		fmt.Printf("len: %d , caps : %d \r\n", len(num1), cap(num1))
	}
}
