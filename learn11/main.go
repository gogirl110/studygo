package main

//二元运算符
import "fmt"

func main() {
	const v = 20 //无显式类型声明的变量

	var a byte = 10
	b := v + a //v自动转换为byte/unit8类型
	fmt.Printf("%T,%v\n", b, b)

	const c float32 = 1.2
	d := c + v //v自动转换为float32类型
	fmt.Printf("%T, %v\n", d, d)
}
