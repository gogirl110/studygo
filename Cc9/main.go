package main

import "fmt"

func main() {
	_, numb, strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, strs)
}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}

//常量是一个简单值的标识符，在程序运行时，不会被修改的量。
//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
