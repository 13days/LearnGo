package main

import (
	"fmt"
	"math"
)

// 函数定义
func funcFirst()  {
	// 开方
	var getSquare = func(x float64) float64{
		return math.Sqrt(x)
	}
	fmt.Println(getSquare(2))
}

// 函数作为参数实现回调
func callBack(x int) int {
	fmt.Println("回调中:", x)
	return x*x
}
func useCallBack(x int, f func(int)int)  {
	y := f(x)
	fmt.Println("回调后:",y)
}
func testCallBack()  {
	useCallBack(1, callBack)
}

func main() {
	//funcFirst()
	testCallBack()
}
