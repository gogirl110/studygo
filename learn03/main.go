package main

/*//数据
//切片slice可实现类似动态数组的功能
func main() {
	x := make([]int, 0, 5)//创建容量为5的切片

	for i := 0; i < 8; i++ {
		x = append(x, i)//追加数据。当超出容量限制时，自动分配更大的存储空间
	}

	fmt.Println(x)
}
*/

/*
//将字典（map）类型内置，可直接从运行时层面获得性能优化
func main() {
	m := make(map[string]int) //创建字典类型对象

	m["a"] = 1 //添加或设置

	x, ok := m["ok"] //使用ok-idiom获取值，可知道key/value是否存在
	fmt.Println(x, ok)

	delete(m, "a") //删除
}
*/

/*
//结构体（struct）可匿名嵌入其他类型
import (
	"fmt"
)

type user struct { //结构体类型
	name string
	age  byte
}

type manager struct {
	user  //匿名嵌入其他类型
	title string
}

func main() {
	var m manager

	m.name = "Tom" //直接访问匿名字段的成员
	m.age = 29
	m.title = "CTO"

	fmt.Println(m)
}
*/

/*//方法
//可以为当前包内的任意类型定义方法
type X int

func (x *X) inc() { //名称前的参数称作receiver，作用类似python self
	*x++
}

func main() {
	var x X
	x.inc()
	println(x)
}
*/

/*
//还可以直接调用匿名字段的方法，这种方式可实现与继承类似的功能
import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	m.name = "Tom"
	m.age = 29

	println(m.ToString()) //调用user.ToString()
}
*/

//接口
//采用duck type方式，无须在实现类型上添加显示声明
import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface { //接口类型（空接口类型interface{},用途类似OOP里的systen.Object,可接收任意类型对象）
	Print()
}

func main() {
	var u user
	u.name = "Tom"
	u.age = 29

	var p Printer = u //只要包含接口所需的全部方法，即表示实现了该接口
	p.Print()
}
