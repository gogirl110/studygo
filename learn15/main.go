package main

/*
//返回值
//有返回值的函数，必须有明确的return终止语句
func test(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
}
*/

/*
//除非有panic
func test(x int)int  {
	for{
		break
	}
}
*/

/*
func div(x,y int) (int error) {//多返回值列表必须使用括号
	if y==0{
		return 0,errors.New(division by zero)
	}

	return x/y,nil
}
*/

/*
//多返回值可用作其他函数调用实参，或当作结果直接返回
func div(x, y int) (int error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}

	return x / y, nil
}
func log(x int, err error) {
	fmt.Println(x, err)
}

func test() (int, error) {
	return div(5, 0) //多返回值用作return结果
}

func main() {
	log(test()) //多返回值用作实参
}
*/

/*
//命名返回值
func div(x,y int) (z int,err error) {
	if y=0{
		err=errors.New(division by zero)
		return
	}
	z=x/y
	return
}
*/

/*
//匿名函数
//直接执行：
func main() {
	func(s string) {
		println(s)
	}("hello,world!")
}
//赋值给参数：
func main() {
	add := func(x, y int) int {
		return x + y
	}

	println(add(1, 2))
}
//作为参数：
func test(f func()) {
	f()
}

func main() {
	test(func() {
		println("hello,world!")
	})
}
//作为返回值：
func test() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	add := test()
	println(add(1, 2))
}
*/

/*
//普通函数和匿名函数都可作为结构体字段，或经通道传递
func testStruct() {
	type calc struct {
		mul func(x, y int) int
	}

	x := calc{
		mul: func(x, y int) int {
			return x * y
		},
	}

	println(x.mul(2, 3))
}

func testChannel() {
	c := make(chan func(int, int) int, 2)

	c <- func(x, y int) int {
		return x + y
	}

	println((<-c)(1, 2))
}
*/

/*
//闭包
func test(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	f := test(123)
	f()
}
*/

/*
func test(x int) func() {
	println(&x)

	return func() {
		println(&x, x)
	}
}
func main() {
	f := test(0x100)
	f()
}
*/

/*
//延迟求值
func test() []func() {
	var s []func()
	for i := 0; i < 2; i++ {
		s = append(s, func() {
			println(&i, i)
		})
	}

	return s
}

func main() {
	for _, f := range test() {
		f()
	}
}
*/

/*
//解决方法就是每次有用不同的环境变量或传参复制，让各自闭包环境不同
func test() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			println(&x, x)
		})
	}

	return s
}
*/

//多个匿名函数引用同一个环境变量，任何的修改行为都会影响其它函数取值，在并发模式下可能需要同步处理
func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	a, b := test(100)
	a()
	b()
}
