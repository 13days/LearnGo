package main

import "fmt"

const MAX int = 3
func test1(){
	a := [MAX]int{1,2,3}
	var ptr [MAX]*int
	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}
	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
	}
}

// 错误的range访问,应为range是拷贝到一个临时地址中去的,所以地址都是一样的
func test2()  {
	number := [MAX]int{5, 6, 7}
	var ptrs [MAX]*int //指针数组
	//将number数组的值的地址赋给ptrs
	for i, x := range &number {
		ptrs[i] = &x
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}

	fmt.Println("=============正确方式============")
	num := [MAX]int{5, 6, 7}
	var ips [MAX]*int //指针数组
	//将number数组的值的地址赋给ptrs
	for i := 0; i < MAX; i++ {
		ips[i] = &num[i]
	}
	for i, x := range ips {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i,*x, x)
	}
}

func main() {
	//test1()
	test2()
}

