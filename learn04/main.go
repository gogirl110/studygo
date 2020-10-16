package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	fmt.Println(Books{"Go语言", "https://www.runoob.com", "Go语言教程", 6495407})
	fmt.Println(Books{title: "Go语言", author: "https://www.runoob.com", subject: "Go语言教程", book_id: 6495407})
	fmt.Println(Books{title: "Go语言", author: "https://www.runoob.com"})
}
