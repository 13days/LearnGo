package main

import "fmt"

/**
函数的本质是变量
不加括号指向函数地址
加括号表示函数调用
 */
func f(a,b int)  {
	fmt.Println(a,b)
}
func main() {
	// 打印函数地址
	fmt.Println(f)
	var a func(int,int)
	a = f
	a(10,12)
	f(10,12)
}
