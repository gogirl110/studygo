package main

//定义 var关键字
/*
var x int     //自动初始化为0
var y = false //自动推断为bool类型
*/
/*
func main() {
	var x, y int      //相同类型的多个变量
	a, s := 100, "abc"//不同类型初始化值

	println(x, y, a, s)
}
*/

/*
//简短模式
func main() {
	x := 100
	a, s := 1, "abc"

	println(x, a, s)
}
*/

//简短模式限制：1定义变量，同时显式初始化 2不能提供数据类型 3只能用在函数内部
/*
var x = 100

func main() {
	println(&x, x) //全局变量

	x := "abc"
	println(&x, x) //重新定义和初始化同名局部变量
}
*/

/*
//简短模式并不总是重新定义变量，也可能是部分退化的赋值操作
func main() {
	x := 100
	println(&x)

	x, y := 200, "abc" //x退化为赋值操作，仅有y是变量定义

	println(&x, x)
	println(y)
}
*/

/*
//退化赋值的前提条件是:最少有一个新变量被定义，且必须是同一作用域
func main() {
	x := 100
	println(&x, x)

	{
		x, y := 200, 300 //不同作用域，全部是新变量定义
		println(&x, x, y)
	}
}
*/

/*
//在处理函数错误返回值时，退化赋值允许我们重复使用err变量
import (
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	println(f, err)

	buf := make([]byte, 1024)
	n, err := f.Read(buf) //err退化赋值，n新定义
	println(n, err)
}
*/

//多变量赋值
func main() {
	x, y := 1, 2
	x, y = y+3, x+2

	println(x, y)
}
