package main

import "fmt"

// 定义一种类型
type MyFloat float64

// 方法
func (n *MyFloat)add()  {
	*n = *n + 1
}
// 方法
func (n *MyFloat)show()  {
	fmt.Println(*n)
}

func main() {
	var a = MyFloat(1)
	a.add()
	a.show()

	b := MyFloat(2)
	b.show()

	var c MyFloat = 3
	c.show()
}