package main

/*
func main()  {
	var c int             //c代替count
	for i :=0;i<10;i++{   //i代替index
		c++
	}

	println(c)
}
*/

/*
//空标识符"_"为特殊成员，通常作为忽略占位符使用，可作为表达式左值，无法读取内容
import "strconv"

func main()  {
	x,_ :=strconv.Atoi("12")//忽略Atoi的err返回值
	println(x)
}
//空标识符可用来临时规避编译器对未使用变量和导入包的错误检查，但是注意，它是预置成员，不能重新定义
*/


/*
func main()  {
	const x =123
	println(x)

	const y=1.23//未使用，不会引发编译错误

	{
		const x="abc"//在不同作用域定义同名常量
		println(x)
	}
}	
*/

/*
//显式指定类型，必须确保常量左右值类型一致，需要时可做显式转换。右值不能超出常量类型取值范围，否则会引发溢出错误
const(
	x,y int =99,-999 //x被指定为int类型，须显式转换为byte类型
	b byte  =byte(x) //错误:constant -999 overflows uint8
	n       =uint8(y)
)
*/

/*
//常量值也可以是某些编译器能计算出结果的表达式，如unsafe.Sizeof、len、cap等
import"unsafe"

const(
	ptrSize=unsafe.Sizeof(uintptr(0))
	strSize=len("hello,word!")
)
*/

//在常量组中如不指定类型和初始化值，则与上一行非空常量右值（表达式文本）相同
import "fmt"

func main()  {
	const(
		x uint16=120 
		y              //与上一行x类型、右值相同
		s ="abc"
		z              //与s类型、右值相同
	)

	fmt.Printf("%T,%v\n",y,y)//输出类型和值
	fmt.Printf("%T,%v\n",z,z)
}