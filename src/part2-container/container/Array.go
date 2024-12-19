package container

import (
	"fmt"
	"strconv"
)

/**
 * 数组
 * @author ch
 * @date 2024/12/12 19:31
 */
func main2() {
	// 数组定义[num]时 是不同类型的数组 而不是不同数量的同类型数组
	var num [3]int
	num[0] = 1
	fmt.Println(num)

	// 数组常用操作
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	// for
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// for range i : index  x : value
	for i, v := range a {
		fmt.Println(i, v)
	}

	// 初始化
	var b = [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b, c)

	// 只初始化某位置
	d := [5]string{1: "a", 3: "c"}
	fmt.Println(d)

	// 省略
	f := [...]string{1: "a", 4: "c"}
	fmt.Println(f)

	// 数组间的比较 需要数组长度也就是类型相同才可以进行比较，自动判断内容
	if d == f {
		fmt.Println("d == f")
	} else {
		fmt.Println("d != f")
	}

	// 二维数组
	var g [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			fmt.Println(g[i][j])
		}
	}

	for _, v := range g {
		for _, vv := range v {
			fmt.Print(strconv.Itoa(vv) + " ")
		}
		fmt.Println()
	}

	for _, v := range g {
		fmt.Println(v)
	}
}
