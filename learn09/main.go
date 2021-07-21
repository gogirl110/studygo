package main

/*
//引用类型。特指slice、map、channel这三种预定义类型
func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}
func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func main() {
	m := mkmap()
	println(m["a"])

	s := mkslice()
	println(s[0])
}
//除new/make函数以外，也可以使用初始化表达式，编译器生成的指令基本相同
*/

//new函数也可以为引用类型分配内存，但是不完整创建
import "fmt"

func main() {
	p := new(map[string]int) //函数返回指针
	m := *p
	m["a"] = 1
	fmt.Println(m)
}
