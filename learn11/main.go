package main

/*
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
*/

/*
//位移右操作数必须是无符号整数，或可以转换的无显式类型常量
func main() {
	b := 23 //b是有符号int类型变量
	x := 1 << b
	println(x)
}
*/

/*
//如果是非常量位移表达式，那么会优先将无显式类型的常量左操作数转型
func main() {
	a := 1.0 << 3               //常量表达式（包括常量展开）
	fmt.Printf("%T,%v\n", a, a) //int，8

	var s uint = 3
	b := 1.0 << s
	fmt.Printf("%T,%v\n", b, b)

	var c int32 = 1.0 << s      //自动将1.0转换int32类型
	fmt.Printf("%T,%v\n", c, c) //int32，8
}
*/

/*
//位运算符
const (
	read byte = 1 << iota
	write
	exec
	freeze
)

func main() {
	a := read | write | freeze
	b := read | freeze | exec
	c := a &^ b //相当于a^read^freeze,但不包括exec

	fmt.Printf("%04b &^ %04b =%04b\n", a, b, c)
}
*/

/*
//自增
func main()  {
	a :=1
	++a

	if (a++)>1 {
	}

	p:=&a
	*p++  //相当于(*p)++
	println(a)
}
*/

/*
func main() {
	x := 10

	var p *int = &x //获取地址，保存到指针变量
	*p += 20        //用指针间接引用，并更新对象

	println(p, *p) //输出指针所存储的地址，以及目标对象
}
*/

/*
func main() {
	m := map[string]int{"a": 1}
	println(&m["a"])
}
*/

/*
func main() {
	x := 10
	p := &x

	p++
	var p2 *int = p + 1

	p2 = &x
	println(p == p2)
}
*/

/*
//零长度（zero-size）对象的地址是否相等和具体的实现版本有关，不过肯定不等于nil
func main() {
	var a, b struct{}

	println(&a, &b)
	println(&a == &b, &a == nil)
}
*/
