package main

//import "fmt"

func main() {
	/*
		fmt.Println("hello world !")

		//变量
		var x int32
		var s = "hello,world!"
		println(x, s)

		//省略var
		x := 100
		println(x)
	*/

	//表达式
	//1,if
	/*
		x := 100

		if x > 0 {
			println("x")
		} else if x < 0 {
			println("-x")
		} else {
			println("0")
		}
	*/
	//2,switch
	/*
		x := 100
		switch {
		case x > 0:
			println("x")
		case x < 0:
			println("-x")
		default:
			println("0")
		}
	*/
	//3,for.1
	/*
		for i := 0; i < 5; i++ {
			println(i)
		}
		for i := 4; i >= 0; i-- {
			println(i)
		}
	*/
	//3,for.2
	/*
		x := 0
		for x < 5 { //相对于while（x<5){...}
			println(x)
			x++
		}
	*/
	//3,for.3
	/*
		x := 4
		for { //相当于while（true）{...}
			println(x)
			x--

			if x < 0 {
				break
			}
		}
	*/
	//3,for.4
	//在迭代遍历时，for...range除元素外，还可以返回索引
	x := []int{100, 101, 102}

	for i, n := range x {
		println(i, ":", n)
	}
}
