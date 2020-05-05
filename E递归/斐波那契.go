package main

import "fmt"

// 斐波那契数列
func fib(n int) int {
	if n<=2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(6))
}
