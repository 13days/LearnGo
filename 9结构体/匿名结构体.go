package main

import "fmt"

func main() {
	// 匿名结构体
	a := struct {
		name string
		age int
	}{
		name:"23",
		age:1,
	}
	fmt.Println(a)
}
