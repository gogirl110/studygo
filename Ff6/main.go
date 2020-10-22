package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

//如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 close() 函数来关闭
