package main

/*
//swich
//与if类似，switch语句也用于选择执行，但具体使用场景会有所不同
func main() {
	a, b, c, x := 1, 2, 3, 2

	switch x { //将想与case条件匹配
	case a, b: //多个匹配条件命中其一即可(OR)，变量
		println("a|b")
	case c: //单个匹配条件
		println("c")
	case 4: //常量
		println("d")
	default:
		println("z")
	}
}
*/

/*
func main() {
	switch x := 5; x {
	default: //编译器确保不会先执行default块
		x += 100
		println(x)
	case 5:
		x += 50
		println(x)
	}
}
//考虑到default作用类似else，建议将其放置switch末尾
*/

/*
//相邻的空case不构成多条件匹配
switch x {
case a:
case b:
	println("b")
}
*/

/*
//不能出现重复的case常量值
func main() {
	switch x := 5; x {
	case 5:
		println("a")
	case 6, 5:
		println("b")
	}
}
*/

/*
//无须显式执行break语句，case执行完毕后自动中断
func main() {
	switch x := 5; x {
	default:
		println(x)
	case 5:
		x += 10
		println(x)

		fallthrough //继续执行下一case，但不再匹配条件表达式
	case 6:
		x += 20
		println(x)

		//fallthrough
	}
}
*/

/*
//fallthrough必须放在case块结尾，可使用break语句阻止
func main() {
	switch x := 5; x {
	case 5:
		x += 10
		println(x)

		if x >= 15 {
			break //终止，不再执行后续语句
		}

		fallthrough //必须是case块的最后一条语句
	case 6:
		x += 20
		println(x)
	}
}
*/

/*
func main() {
	switch x := 5; { //相当于"switch x := 5;true{...}"
	case x > 5:
		println("a")
	case x > 0 && x <= 5:
		println("b")
	default:
		println("z")
	}
}
*/

/*
//for
func count() int {
	print("count.")
	return 3
}

func main() {
	for i, c := 0, count(); i < c; i++ { //初始化语句的count函数仅执行一次
		println("a", i)
	}

	c := 0
	for c < count() { //条件表达式中的count重复执行
		println("b", c)
		c++
	}
}
*/

/*
//可用for...range完成数据迭代，支持字符串、数组、数组指针、切片、字典、通道类型，返回索引、键值数据
func main() {
	data := [3]string{"a", "b", "c"}

	for i, s := range data {
		println(i, s)
	}
}
*/

/*
//允许返回单值，或用"_"忽略
func main() {
	data := [3]string{"a", "b", "c"}

	for i := range data { //只返回1st value
		println(i, data[i])
	}

	for _, s := range data { //忽略1st value
		println(s)
	}

	for range data { //仅迭代，不返回。可用来执行清空channel等操作
	}
}
*/

/*
//无论普通for循环，还是染个迭代，其定义的局部变量都会重复使用
func main() {
	data := [3]string{"a", "b", "c"}

	for i, s := range data {
		println(&i, &s)
	}
}
*/

/*
func main() {
	data := [3]int{10, 20, 30}

	for i, x := range data { //从data复制品中取值
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}

		fmt.Printf("x:%d,data:%d\n", x, data[i])
	}

	for i, x := range data[:] { //仅复制slice，不包括底层array
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}

		fmt.Printf("x:%d,data:%d\n", x, data[i])
	}
}
*/

/*
//如果range目标表达式是函数调用，也仅被执行一次
func data() []int {
	println("origin data.")
	return []int{10, 20, 30}
}

func main() {
	for i, x := range data() {
		println(i, x)
	}
}
*/

/*
//goto,countinue,break
//使用goto前，须先定义标签。标签区分大小写，且使用的标签会引发编译错误
func main() {
start: //错误:label start defrined and not used
	for i := 0; i < 3; i++ {
		println(i)

		if i > 1 {
			goto exit
		}
	}

exit:
	println("exit.")
}
*/

/*
//不能跳转到其它函数，或内层代码块内
func test() {
test:
	println("test")
	println("test exit")
}

func main() {
	for i := 0; i < 3; i++ {
	loop:
		println(i)
	}

	goto test
	goto loop
}
*/

/*
//和goto定点跳转不同，break、coutinue用于中断代码块执行
//break:用于switch、for、select语句，终止整个代码块执行
//continue：仅用于for循环，终止后续逻辑，立即进入下一轮循环
func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //立即进入下一轮循环
		}
		if i > 5 {
			break //立即终止整个for循环
		}

		println(i)
	}
}
*/

//配合标签，break和continue可在多层嵌套指定目标层级
func main() {
outer:
	for x := 0; x < 5; x++ {
		for y := 0; y < 10; y++ {
			if y > 2 {
				println()
				continue outer
			}

			if x > 2 {
				break outer
			}

			print(x, ":", y, " ")
		}
	}
}
