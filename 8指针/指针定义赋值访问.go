package main

import "fmt"

func testPtr()  {
	var a int = 20
	var ip *int // 声明指针
	ip = &a // 赋值
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
}

// 空指针
func testNil()  {
	var ptr *int
	fmt.Printf("ptr的值为:%x\n", ptr)
	//fmt.Printf("*ptr的值为:%x\n", *ptr) // 异常
	if ptr == nil{
		fmt.Println("空指针")
	}else{
		fmt.Println("不是空指针")
	}
}
func main() {
	//testPtr()
	testNil()
}

