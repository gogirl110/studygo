package main

/*
import (
	"fmt"
	"math"
)

func main() {
	a, b, c := 100, 0144, 0x64

	fmt.Println(a, b, c)
	fmt.Printf("0b%b,%#o,%#x\n", a, a, a)

	fmt.Println(math.MinInt8, math.MaxInt8)
}
//支持八进制、十六进制以及科学计数法。标准库math定义了各数字类型的取值范围
*/

/*
//标准库strconv可在不同进制（字符串）间转换
import "strconv"

func main() {
	a, _ := strconv.ParseInt("1100100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	c, _ := strconv.ParseInt("64", 16, 32)

	println(a, b, c)

	println("0b" + strconv.FormatInt(a, 2))
	println("0" + strconv.FormatInt(a, 8))
	println("0x" + strconv.FormatInt(a, 16))
}
*/

/*
//使用浮点数时，须注意小数位的有效精度
func main() {
	var a float32 = 1.1234567899 //注意：默认浮点类型是float64
	var b float32 = 1.12345678
	var c float32 = 1.123456781

	println(a, b, c)
	println(a == b, a == c)
	fmt.Printf("%v %v, %v\n", a, b, c)
}
*/

//别名
byte    alias for uint8
rune    alias for int32

//别名无须转换，可直接赋值
func test(x byte) {
	println(x)
}

func main() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b

	test(c)
}
