package main

/**
 * 基础数据类型
 * @date 2024/12/12 11:24
 */
//func main() {
//	// int8 - 64 uint8-64 无符号数为有符号数范围的一倍 因为二进制中首位代表正负
//	// 字符 byte (uint8), rune (uint32)中英文字符时使用
//	// 浮点 float32 64
//	// 字符串 string
//	// 数值类型转换 使用类型强转
//
//	// strconv主要转换字符串类型
//	itoa := strconv.Itoa(12)      // int转换str
//	atoi, _ := strconv.Atoi("15") // 字符串转换int类型
//	fmt.Println(itoa, atoi)
//	// 使用 parse 进行转换 str - > 基础类型
//	fmt.Printf("%d", 12)
//	float, err := strconv.ParseFloat("12.5", 64)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(float)
//	}
//	// 使用format 将基础类型转换为 str
//	formatBool := strconv.FormatBool(true)
//	fmt.Println("转换布尔：" + formatBool)
//
//}
