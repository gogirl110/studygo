package main

import "fmt"

func main() {
	source := []int{7, 6, 10, 111, 2000000, 1}
	var tmp int
	fmt.Println("未排序前:", source)
	for i := 0; i < len(source)-1; i++ {
		fmt.Println("第", i+1, "次冒泡", source)
		for i := 0; i < len(source)-1-i; i++ {
			if source[i] > source[i+1] {
				tmp = source[i+1]
				source[i+1] = source[i]
				source[i] = tmp
			}
		}
	}
	fmt.Println("最终结果", source)
}
