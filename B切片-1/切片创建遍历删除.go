package main

import "fmt"

// 空值测试
func testNil()  {
	var numbers []int

	PrintSlice(numbers)

	if(numbers == nil){
		fmt.Printf("切片是空的")
	}
}

// 创建切片
func f1()  {
	// 长度3,最大容量5,[]里是不写长度的,写长度的是数组,数组是定长的不可扩容,切片可以
	// 默认长度的len,容量是cap,append到超过cap会扩容
	var numbers = make([]int,3,5)
	PrintSlice(numbers)
	testNil()
}

// 切片复制和追加
func f2() {
	var numbers []int
	numbers = append(numbers, 1, 3, 4, 6)
	PrintSlice(numbers)

	// 切片复制
	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	copy(numbers1, numbers)
	//继续添加
	numbers1 = append(numbers1,1,2)
	PrintSlice(numbers1)
}

// 切片删除
func f3()  {
	var numbers = []int{1,2,4,5}
	// 删除位置1
	// numbers[2:len(numbers)]... 表示把后段的所有位置添加到前段中
	numbers = append(numbers[0:1], numbers[2:len(numbers)]...)
	PrintSlice(numbers)
}
func main()  {
	//f1()
	//f2()
	f3()
}

// 打印
func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}