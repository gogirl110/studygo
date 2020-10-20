package main

import "fmt"

func main() {
	var numbers []int

	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

//一个切片在未初始化之前默认为 nil，长度为 0
