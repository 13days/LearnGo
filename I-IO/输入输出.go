package main

import "fmt"

/**
%T 类型
%b 2进制
%o 8进制
%x 16进制 a-b-c
%X 16进制 A-B-C
$p 内存地址
%c

输入:

 */
func main() {
	var a int
	var b float64
	fmt.Scanln(&a,&b)
	fmt.Println(a,b)
}
