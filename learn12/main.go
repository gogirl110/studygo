package main

import (
	"errors"
	"log"
	"strconv"
)

/*
//初始化
type data struct{
	x int
	s string
}

var a data=data{1,"abc"}

b:=data{
	1,
	"abc",
}

c:=[]int{
	1,
	2}

d:=[]int{1,2,
	3,4,
	5,
}
*/

/*
//流控制
//if...else...
func main() {
	x := 3

	if x > 5 {
		println("a")
	} else if x < 5 && x > 0 {
		println("b")
	} else {
		println("z")
	}
}
*/

/*
//局部变量的有效范围包含整个if/else
func main() {
	x := 10

	if xinit(); x == 0 { //优先执行xinit函数
		println("a")
	}
	if a, b := x+1, x+10; a < b { //定义一个或多个局部变量(也可以是函数返回值)
		println(a)
	} else {
		println(b)
	}
}
*/

/*
func main() {
	x := 8

	if x > 5 { //优先判断，条件表达式结果为true
		println(a)
	} else if x > 7 {
		println("b")
	}
}
*/

/*
//尽可能减少代码块嵌套，让正常逻辑处于相同层次
func check(x int) error {
	if x <= 0 {
		return errors.New("x<=0")
	}

	return nil
}

func main() {
	x := 10

	if err := check(x); err == nil {
		x++
		println(x)
	} else {
		log.Fatalln(err)
	}
}
*/

/*
//基于重构原则，我们应该保持代码块功能的单一性
func main() {
	x := 10

	if err := check(x); err != nil {
		log.Fatalln(err)
	}

	x++
	println(x)
}
*/

/*
func main() {
	s := "9"

	n, err := strconv.ParseInt(s, 10, 64) //使用外部变量

	if err != nil {
		log.Fatalln(err)
	} else if n < 0 || n > 10 { //也可考虑拆分另一个独立if块
		log.Fatalln("invalid number")
	}

	println(n) //避免if局部变量将该逻辑放到else块
}
*/

/*
//对于某些过于复杂的组合条件，建议将其重构为函数
func main() {
	s := "9"

	if n, err := strconv.ParseInt(s, 10, 64); err != nil || n < 0 || n > 10 || n%2 != 0 {
		log.Fatalln("invalid number")
	}

	println("ok")
}
*/

//条件语句独立更易于测试，同样会改善代码可维护性
func check(s string) error {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil || n < 0 || n > 10 || n%2 != 0 {
		return errors.New("invalid number")
	}

	return nil
}

func main() {
	s := "9"

	if err := check(s); err != nil {
		log.Fatalln(err)
	}

	println("ok")
}
