package main

import (
	"fmt"
)

func main()  {
	const (
		a = iota 
		b = iota
		c = iota
	)
	fmt.Printf("a = %d,b = %d,c = %d\n", a, b, c)

	const d = iota
	fmt.Printf("d = %d\n", d)

	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 =%d\n, c1 = %d\n",a1,b1,c1)
	
	const
}