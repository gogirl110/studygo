package main

import "fmt"

/*
//函数
//定义
func hello() {
	println("hello,world!")
}

func exec(f func()) {
	f()
}

func main() {
	f := hello
	exec(f)
}
*/

/*
//从函数返回局部变量指针是安全的，编译器会通过逃逸分析(escape analysis)来决定是否在堆上分配内存
func test() *int {
	a := 0x100
	return &a
}

func main() {
	var a *int = test()
	println(a, *a)
}
*/

/*
//参数
//在参数列表中，相邻的同类型参数可合并
func test(x,y int,s string,_bool)*int {
	return nil
}

func main()  {
	test(1,2,"abc")
}
*/

/*
//参数可视作函数局部变量，因此不能在相同层次定义同名变量
func add(x, y int) int {
	x := 100
	var y int
	return x + y
}
*/

/*
func test(x *int) {
	fmt.Printf("pointer:%p,tanger:%v\n", &x, x) //输出形参x的地址
}

func main() {
	a := 0x100
	p := &a
	fmt.Printf("pointer:%p,target:%v\n", &p, p) //输出实参p的地址

	test(p)
}
*/

/*
//指针参数导致实参变量被分配到堆上
func test(p *int) {
	go func() { //延长p生命周期
		println(p)
	}()
}

func main() {
	x := 100
	p := &x
	test(p)
}
*/

/*
//要实现传出参数(out)，通常建议使用返回值。当然，也可以继续用二级指针
func test(p **int) {
	x := 100
	*p = &x
}

func main() {
	var p *int
	test(&p)
	println(*p)
}
*/

/*
//如果函数参数过多，建议将其重构为一个复合结构类型，也算是变相实现可选参数和命名实参功能
type serverOption struct {
	addresss string
	port     int
	path     string
	timeout  time.Duration
	log      *log.Logger
}

func newOption() *serverOption {
	return &serverOption{ //默认参数
		addresss: "0.0.0.0",
		port:     8080,
		path:     "/var/test",
		timeout:  time.Second * 5,
		log:      nil,
	}
}

func server(option *serverOption) {}

func main() {
	opt := newOption()
	opt.port = 8085 //命名参数设置

	server(opt)
}
*/

/*
//变参
//变参本质上就是一个切片。只能接收一到多个同类型参数，且必须放在列表尾部
func test(s string, a ...int) {
	fmt.Printf("%T,%v\n", a, a)
}

func main() {
	test("abc", 1, 2, 3, 4)
}
*/

/*
//将切片作为变参时，须进行展开操作。如果是数组，先将其转换为切片
func test(a ...int) {
	fmt.Println(a)
}

func main() {
	a := [3]int{10, 20, 30}
	test(a[:]...) //转换为slice后展开
}
*/

func test(a ...int) {
	for i := range a {
		a[i] += 100
	}
}

func main() {
	a := []int{10, 20, 30}
	test(a...)

	fmt.Println(a)
}
