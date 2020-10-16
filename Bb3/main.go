package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d,b = %d,c = %d\n", a, b, c)

	const d = iota
	fmt.Printf("d = %d\n", d)

	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 =%d\n, c1 = %d\n", a1, b1, c1)

}

// 现在这个代码有bug 编译不过去， 你自己处理一下
// 现在可以了

//如果左边显示哦红色 一般情况下 指的是 代码有bug
// 你现在操作Bb2 把bug fix掉
