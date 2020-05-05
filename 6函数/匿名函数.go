package main

import "fmt"

func main() {
	func(){
		fmt.Println("匿名函数1")
	}()// 调用

	c := func(a,b int)int{
		return a+b
	}(1,2)
	fmt.Println(c)
}
