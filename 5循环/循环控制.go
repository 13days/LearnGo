package main

import "fmt"

// break
func test5()  {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++;
		if a > 15 {
			/* 使用 break 语句跳出循环 */
			break
		}
	}
}

// 标记break
func test6()  {
	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("  i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
	re: // 默认是退出当前循环,可以使用标记,指定退出具体的循环层
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("  i2: %d\n", i2)
			break re
		}
	}
}

// continue
func test7()  {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}

// 使用标记的continue
func test8()  {
	// 不使用标记
	fmt.Println("---- continue ---- ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("  i2: %d\n", i2)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- continue label ----")
	re: for i := 1; i <= 3; i++ { // continu本标记下的循环
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("  i2: %d\n", i2)
			continue re
		}
	}
}

// goto
func test9()  {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
	LOOP: for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
func main() {
	test9()
}
