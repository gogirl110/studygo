package main

import "fmt"

/*
//类型转换
func main() {
	x := 100
	//不能将非bool类型结果当作true/false使用
	var b bool = x

	if x {

	}
}
*/

/*语法歧义
func main() {
	x := 100
	p := *int(&x)

	println(p)
}
*/

/*
//自定义类型
type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func main() {
	f := read | exec
	fmt.Printf("%b\n", f)
}
*/

/*
//和var、constant类似，多个type定义可合并成组，可在函数或代码块内定义局部类型
func main() {
	type (
		user struct {
			name string
			age  uint8
		}

		event func(string) bool
	)

	u := user{"Tom", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")
}
*/

/*
func main()  {
	type data int
	var d data=10

	var x int =d
	println(x)

	println(d==x)
}
*/

/*未命名参数顺序
func main()  {
	var a struct{
		x  int 'x'
		s string 's'
	}

	var b struct{
		x int
		s string
	}

	b=a

	fmt.Println(b)
}
*/

/*
//函数的参数顺序也属于签名组成部分
func main() {
	var a func(int, string)
	var b func(string, int)

	b = a

	b("s", 1)
}
*/

func main() {
	type data [2]int
	var d data = [2]int{1, 2} //基础类型相同，右值为未命名类型

	fmt.Println(d)

	a := make(chan int, 2)
	var b chan<- int = a //双向通道转换为单向通道，其中b为未命名类型

	b <- 2
}
