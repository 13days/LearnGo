package main

import "fmt"

// 普通循环
func test1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 省略
func test2()  {
	sum := 1
	i := 1
	for sum < 50{
		sum += i
		i++
	}
	fmt.Println(sum)
}

// 无限循环
func test3()  {
	sum,i:= 1,1
	for{
		sum += i
		i++
		if sum>50{
			break
		}
	}
	fmt.Println(sum)
}

// for-each
func test4() {
	strings := []string{"root", "usr"}
	for i, s := range strings{
		fmt.Println(i,s)
	}

	numbers := []int{1,2,3,4}
	for i,num := range numbers{
		fmt.Println(i,num)
	}
}
func main() {
	test4()
}
