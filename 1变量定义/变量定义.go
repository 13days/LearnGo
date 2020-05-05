package main

import "fmt"

// 如何声明变量
func test0()  {
	var a string = "root"
	fmt.Println(a)
	var x,y,z int = 1, 2, 3
	fmt.Println(x,y,z)
}
// 初始化
func test1() {
	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)
}
// 以下几种类型为 nil：
func test2() {
	var a *int
	var b []int
	var c map[string] int
	var d chan int
	var e func(string) int
	var f error // error 是接口
	fmt.Println(a,b,c,d,e,f)
}

func test3() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// 自动推断变量类型
func test4() {
	var d = true
	fmt.Println(d)
}
// 省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
func test5() {
	// 报错
	//var a int
	//a := 1
	//fmt.Println(a)

	a := 1
	fmt.Println(a)

	b,c := 1,2
	fmt.Println(b,c)

}

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func test6() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
func main()  {
	//test0()
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
	test6()
}
