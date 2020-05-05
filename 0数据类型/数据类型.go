package main

import "fmt"

/*
基本数据类型
	bool
	Number Type:
		int8, int16, int32,	int64, int(电脑几位就几位)
		uint8, uint16, uint32, uint64, uint
		float32, float64
		complex64, complex128
		byte:uint8
		rune:int32
	string

复合数据类型
	array, slice, map, function, pointer, struct, interface, channel...

 */

func main() {
	a := 100
	fmt.Printf("type=%T,num=%d\n", a, a)

	var x int = 444
	var y int64
	// y = x // cannot use x (type int) as type int64 in assignment
	fmt.Printf("type=%T,num=%d\n", y, x)
}