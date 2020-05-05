package main

import "fmt"

// 通道创建
func testChannel()  {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// range要关闭通道
func fib(n int, c chan int)  {
	defer close(c)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		c <- y
	}
}
func testClose()  {
	c := make(chan int, 10)
	go fib(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c{
		fmt.Println(i)
	}
}
func main() {
	//testChannel()
	testClose()
}
