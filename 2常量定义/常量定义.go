package main

import (
	"fmt"
	"unsafe"
)

// 常量定义
func test1() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}

// 枚举定义
// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
	d = 123
)

func test2() {
	println(a,b,c,d)
}
// iota，特殊常量，可以认为是一个可以被编译器修改的常量。

// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
func test3() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}

// iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），
// 即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
func test4() {
	const (
		i=1<<iota
		j=3<<iota
		k
		l
	)
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}
func main()  {
	// test1()
	// test2()
	// test3()
	test4()
}
