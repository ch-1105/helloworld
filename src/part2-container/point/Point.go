package main

import "fmt"

type Point struct {
	name string
}

func updateObj(point *Point) {
	point.name = "hello"
}

func swap(n1, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}
func mainPoint() {
	// go指针可以直接赋值，底层有优化，但不能进行运算
	// 初始化推荐使用new关键字
	var p = Point{
		name: "world",
	}
	fmt.Println(p.name)
	updateObj(&p)
	fmt.Println(p.name)

	var p1 *Point
	p1 = new(Point)

	var p2 *Point
	p2 = &p

	var p3 = &p2
	fmt.Println(p1, p2, p3)

	// 指针swap
	var a, b = 1, 2
	swap(&a, &b)
	fmt.Println(a, b) //为什么a = 1 , b = 2  swap(){n1, n2 = n2, n1}
	//得修改成 swap(){
	//	temp := *n1
	//	*n1 = *n2
	//	*n2 = temp}

}
