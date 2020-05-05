package main

import "fmt"

type A interface {

}

type B interface {
	A
	name()
}

type C interface {
	A
	do()
}
type BINS struct {
	B
}

func (b BINS)name()  {
	fmt.Println("我叫BINS")
}

type CINS struct {
	C
}
func (c CINS)do()  {
	fmt.Println("CINS do...")
}
func testA(a A)  {
	if ins, ok := a.(B); ok{
		fmt.Println("是B")
		ins.name()
	}else if ins, ok := a.(C); ok{
		fmt.Println("是C")
		ins.do()
	}
}

func testAA(a A)  {
	switch ins := a.(type) {
	case B:
		ins.name()
	case C:
		ins.do()
	}
}
func main() {
	b := BINS{}
	c := CINS{}
	testA(b)
	testA(c)

	testAA(b)
	testAA(c)
}