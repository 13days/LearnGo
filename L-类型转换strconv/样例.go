package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1234"
	// 进制,位数
	b,err := strconv.ParseInt(a, 10, 32)
	fmt.Printf("%T,%s\n", b,err)
	if err==nil{
		fmt.Println("解析成功",b)
	}

	x := true
	str := strconv.FormatBool(x)
	fmt.Println(str)


	i,err := strconv.Atoi("12334")
	fmt.Printf("%T\n", i)
	str = strconv.Itoa(i)
	fmt.Printf("%T\n",str)
}
