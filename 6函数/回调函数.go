package main

import "fmt"

func myOpt(a,b int,f func(int,int) int) int{
	fmt.Println("准备进行调用")
	return f(a,b)
}
func add(a,b int)int{
	fmt.Println("执行加法")
	return a + b
}
func sub(a,b int)int{
	fmt.Println("执行减法")
	return a-b
}
func main() {
	myOpt(1,2,add)
	myOpt(5,1,sub)
}
