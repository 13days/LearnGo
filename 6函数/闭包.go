package main

import "fmt"

// 闭包内联,返回一个匿名函数
func Add(x int) func() int{
	return func()int{
		x++
		return x
	}
}
func testAdd()  {
	testFunc := Add(0)
	fmt.Println(testFunc())
	fmt.Println(testFunc())
	fmt.Println(testFunc())

	testFunc1 := Add(1)
	fmt.Println(testFunc1())
	fmt.Println(testFunc1())
	fmt.Println(testFunc1())
}
func main() {
	testAdd()
}
