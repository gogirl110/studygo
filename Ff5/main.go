package main

import "fmt"

func main() {
	//缓冲区大小为2
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	//获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小
