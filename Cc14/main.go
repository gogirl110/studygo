package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b //相加
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b //相减
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b //相乘
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b //相除
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b //求余
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++ //自增
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--    //自减
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}
