package main

import "fmt"

// 自定义函数
func max(a,b int) int {
	if a > b{
		return a
	}else{
		return b
	}
}

// 多返回值
func swapInt(a,b int) (int, int) {
	return b,a
}
func swapStr(a,b string)(string, string) {
	return b,a
}

// 值传递,复制一份到函数里
func swapWithZhi(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}
func testSwap()  {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a )
	fmt.Printf("交换前 b 的值为 : %d\n", b )

	/* 通过调用函数来交换值 */
	swapWithZhi(a, b)

	fmt.Printf("交换后 a 的值 : %d\n", a )
	fmt.Printf("交换后 b 的值 : %d\n", b )
}

// 引用传递
func swapWithYinYong(x, y *int) {
	var temp int
	temp = *x // 解引用
	*x = *y // 把y的值附到x
	*y = temp // 把temp赋值到y
}
func testSwap1()  {
	a, b := 1, 2
	fmt.Println("交换前:", a, b)
	swapWithYinYong(&a, &b)
	fmt.Println("交换后:", a, b)
}

// 不定长参数
func test(a... int)  {
	for i,x := range a{
		fmt.Println(i,x)
	}
}

// 命名返回值
func f(x,y int) (sum int){
	sum = x + y
	return sum
}
func main() {
	//fmt.Println(max(10,88))
	//fmt.Println(swapInt(10,88))
	//fmt.Println(swapStr("123","321"))
	//testSwap()
	//testSwap1()
	//fmt.Println(f(1,2))
	test([]int{1,2,3,4}...)
	test(1,2,3,4)
}
