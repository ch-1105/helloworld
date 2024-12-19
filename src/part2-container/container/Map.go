package container

import "fmt"

func main4() {
	/* 使用 make 函数 */
	// map_variable := make(map[KeyType]ValueType, initialCapacity)
	map1 := make(map[string]string, 10)
	map2 := map[string]string{
		"c": "c",
		"d": "d",
	}
	map1["a"] = "1"
	map1["b"] = "2"
	fmt.Println(map1)

	str, _ := map1["a"]
	fmt.Println(str)

	// 遍历
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	// 删除，删除不存在的元素不会报错
	delete(map1, "a")

	fmt.Println(map1)

	map2["hh"] = "hh"
	fmt.Println(len(map2))

	// 检测值存在
	_, ok := map2["kk"]
	fmt.Println("打印是否存在： ", ok)

	// if用法  赋值并进行判断
	if _, ok := map2["kk"]; !ok {
		fmt.Println("。。。。。")
	}

	// map不是线程安全的
}
