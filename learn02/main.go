package main

/*
//函数可以定义多个返回值，甚至对其命名
import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	a, b := 10, 2
	c, err := div(a, b)

	fmt.Println(c, err)
}
*/

/*
//函数是第一类型，可作为参数或返回值
func test(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	x := 100

	f := test(x)
	f()
}
*/

//用defer定义延迟调用，无论函数是否出错，它都确保结束前被调用
func test(a, b int) {
	defer println("dispose...")

	println(a / b)
}

func main() {
	test(10, 0)
}
