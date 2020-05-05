package main

import "fmt"

func testPanic()  {
	//defer func() {
	//	if msg := recover(); msg!=nil{
	//		fmt.Println(msg,"程序恢复")
	//	}
	//}()
	defer fmt.Println("第一个defer")
	for i := 0; i < 10; i++ {
		if i==5{
			panic("恐慌了,中断了")
		}
		fmt.Println(i)
	}
	// 只有恐慌前的defer会在抛出恐慌前执行
	defer fmt.Println("第二个defer")
}

func main() {
	defer func() {
		if msg := recover(); msg!=nil{
			fmt.Println(msg,"程序恢复")
		}
	}()
	defer fmt.Println("main的defer-1")
	defer fmt.Println("main的defer-2")
	testPanic()
	defer fmt.Println("main的defer-3")
}