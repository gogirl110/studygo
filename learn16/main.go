package main

/*
//延迟调用
func main()  {
	f,err := os.Open("./main.go")
	fi err !=nil{
		log.Fatalln(err)
	}

	defer f.Close()//仅注册，直到main退出前才执行

	...do something...
}
*/

/*
//延迟调用注册的是调用，必须提供执行所需参数（哪怕为空）。参数值在注册时被复制并缓存起来。如对状态敏感，可改用指针或闭包
func main() {
	x, y := 1, 2

	defer func(a int) {
		println("defer x,y=", a, y)
	}(x)
	x += 100
	y += 200
	println(x, y)
}
*/

/*
//多个延迟注册按FILO次序执行
func main() {
	defer println("a")
	defer println("b")
}
*/

/*
//return语句不是ret汇编指令，它会先更新返回值
func test() (z int) {
	defer func() {
		println("defer:", z)
		z += 100
	}()

	return 100
}

func main() {
	println("test:", test())
}
*/

/*
//误用
//延迟调用在函数结束时才会被使用，不合理的使用方式会浪费更多资源，甚至造成逻辑错误
//循环处理多个日志文件，不恰当的defer导致文件关闭时间延长
func main() {
	for i := 0; i < 10000; i++ {
		path := fmt.Sprintf(".log/%d.txt", i)

		f, err := os.Open(path)
		if err != nil {
			log.Panicln(err)
			continue
		}

		//这个关闭操作在main函数结束时才回执行，而不是当前循环中执行
		//这无端延长了逻辑结束时间和f的生命周期，平白多消耗了内存等资源
		defer f.Close()

		... do something...
	}
}
*/

/*
func main() {
	//日志处理算法
	do := func(n int) {
		path := fmt.Sprintf("./log/%d.txt", n)

		f, err := os.Open(path)
		if err != nil {
			log.Panicln(err)
			continue
		}

		//该延迟调用在此匿名函数结束时执行，而非main
		deferf.Close()

		...do something...
	}

	for i := 0; i < 100000; i++ {
		do(i)
	}
}
*/

/*
//性能
var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	defer m.Unlock()
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func Benchmarkdefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}
//性能要求高且压力大的算法，应避免使用延迟调用
*/

/*
//错误处理
//error
//官方推荐的标准做法是返回error状态
func Scanln(a ...interface{}) (n int, err error)

//标准库将error定义为接口类型，以便实现自定义错误类型
type error interface {
	Error() string
}
*/

/*
//标准库提供了相关创建函数，可方便地创建包含简单错误文本的error对象
var errDivByZero = errors.New("divsion by zero")

func div(x, y int) (int error) {
	if y == 0 {
		return 0, errDivByZero
	}
	return x / y, nil
}

func main() {
	z, err := div(5, 0)
	if err == errDivByZero {
		log.Fatalln(err)
	}

	println(z)
}
//应通过错误变量，而非文本内容来判定错误类别
*/

/*
type DivError struct {//自定义类型
	x, y int
}

func (DivError) Error() string {//实现error接口方法
	return "division by zero"
}

func div(x, y int) (int error) {
	if y == 0 {
		return 0, DivError{x, y}//返回自定义错误类型
	}

	return x / y, nil
}

func main() {
	z, err := div(5, 0)

	if err != nil {
		switch e := err.(type) {//根据类型匹配
		case DivError:
			fmt.Println("e,e.x,e.y")
		default:
			fmt.Println(e)
		}

		log.Fatalln(err)
	}

	println(z)
}
*/

/*
//panic,recover
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	panic("i am dead")
	println("exit.")
}
*/

func test() {
	defer println("test.1")
	defer println("test.2")

}
