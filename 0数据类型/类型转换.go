package main

import "fmt"

/**
静态语言:要求定义赋值运算类型必须一致
语法Type(Value)
 */

func main() {
	var a int64 = 18
	var b int32 = int32(a)
	fmt.Println(b)

	var c float64 = 12.23
	a = int64(c)
	fmt.Println(a)
	
}