package main

import (
	"fmt"
	"unsafe"
)

func TestPtr()  {
	var s string = "sss"
	p := &s
	*(p) = "123"

	// 显示声明指针类型
	*(*string)(p) = "23423"
	fmt.Println(s)

	// 强制类型转换
	//num := 5
	//numPointer := &num
	//
	//flnum := (*float32)(numPointer)
	//fmt.Println(flnum)

	num := 5
	numPointer := &num
	var ptr unsafe.Pointer = unsafe.Pointer(numPointer)
	floatPtr := (*float32)(ptr)
	*floatPtr = 3.0
	fmt.Printf("Type=%T,Value=%.2f\n",floatPtr,*floatPtr)

}
func main()  {
	TestPtr()
}
