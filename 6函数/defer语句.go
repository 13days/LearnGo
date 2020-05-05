package main

import "fmt"

// defer 最先被定义最后被执行,在函数要退出的时候执行
// return = 返回值=x  RET x
// defer = 返回值=x  defer  RET x
func myDefer() int{
	x := 1
	fmt.Println("开始")
	fmt.Println("结束了")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer func() {
		x++ // 修改的是x
	}()
	return x // 这是赋值的x
}

func myDefer2()(x int){
	x = 1
	defer func() {
		x++
	}()//这里传参
	return x
}
func myDefer3()(x int){
	x = 1
	defer func(x int) {
		x++
	}(x)//这里传参,是拷贝,原来不变
	return x
}
func myDefer4()(x int){
	x = 1
	defer func(x *int) {
		*x++
	}(&x)//这里传了内存地址
	return x
}


func main() {
	fmt.Println(myDefer())
	fmt.Println(myDefer2())
	fmt.Println(myDefer3())
	fmt.Println(myDefer4())

}
