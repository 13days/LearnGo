package main

import (
	"fmt"
	"sort"
)
type IntSlice []int

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func test() {
	a := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	sort.Sort(IntSlice(a))
	fmt.Println("After sorted: ", a)
}
func main() {
	// 要排序需要满足3个接口,Len,Less,Swap
	a := []int{1,23,5,3,1,3,5,33}
	sort.Ints(a)
	fmt.Println(a)
	test()
}