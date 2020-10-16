package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 int = 10
	var v2 = 10
	v3 := 10
	fmt.Println("v3 type is", reflect.TypeOf(v3))

	var v4 int
	v4 = 2
	fmt.Println("v4 type is", reflect.TypeOf(v4))
	fmt.Println("v1 type is", reflect.TypeOf(v1))
	fmt.Println("v2 type is", reflect.TypeOf(v2))
}
