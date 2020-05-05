package main

import "fmt"

type Book struct {
	name string
	price int
}
type Home struct {
	book Book
	num  int
}

func main() {
	b := Book{"xx", 10}
	// 结构体是值传递
	h := Home{b,10}
	fmt.Println(h)
	h.book.name = "hh"
	fmt.Println(h)
	fmt.Println(b)
}