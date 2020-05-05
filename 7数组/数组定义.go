package main

import "fmt"

// 数组定义和遍历
func define()  {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] += 100
	}

	for i := 0; i < 10; i++ {
		a[i] += i
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Element[%d]=%d\n", i, a[i])
	}
}

// 数组默认定义
func define1()  {
	// 定长赋值
	var balance = [5]float64{1,2,3,4,5}
	for i := 0; i < 5; i++ {
		fmt.Println(balance[i])
	}

	// 不定长赋值
	xixi := [...]float64{1,2,3,4,5}
	for i := 0; i < len(xixi); i++ {
		fmt.Println(xixi[i])
	}
}

// 数组是值类型
func f()  {
	// 切片是引用类型,非固定长度
	var a []int = []int{1,2,3,5}
	b := a
	b[1] = 1000
	fmt.Println(a)
	fmt.Println(b)

	// 数组是值类型, 固定长度的才是数组
	c := [4]int{1,2,3,5}
	d := c
	c[1] = 10000
	fmt.Println(c)
	fmt.Println(d)
}
func main() {
	//define()
	//define1()
	f()
}
