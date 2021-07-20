package main

/*
//枚举
//Go没有明确意义上的enum定义，不过借助iota标识符实现一组自增常量来实现枚举类型
const (
	x = iota //0
	y        //1
	z        //2
)

const (
	_  = iota             //0
	KB = 1 << (10 * iota) //1<<(10*1)
	MB                    //1<<(10*2)
	GB                    //1<<(10*3)
)
*/

/*
//自增作用范围为常量组，可在多常量定义中使用多个iota，它们各自单独计数，只须确保组中每行常量的列数量相同即可
const(
	_,_ =iota,iota*10  //0,0*10
	a,b                //1,1*10
	c,d                //2,2*10
)
*/

/*
//如中断iota自增，则必须显式恢复，且后续自增值按行序递增，而非C enum那般上一取值递增
const(
	a =iota  //0
	b        //1
	c =100   //100
	d        //100(与上一行常量右值表达式相同)
	e =iota  //4(恢复iota自增，计数包括c、d)
	f        //5
)
*/

/*
//自增默认数据类型为int，可显式指定类型
const(
	a =iota
	b  float32=iota
	c =iota
)
*/

/*
//实际编码中，建议使用自定义类型实现用途明确的枚举类型，但这并不能将取值范围限定在预定义的枚举值内
type color byte//自定义类型

const (
	biack color = iota
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	test(red)
	test(100) //100并未超出color/byte类型取值范围

	x := 2
	test(x)
}
*/

//展开
