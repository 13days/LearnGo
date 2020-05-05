package main

import "fmt"

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}
type ClickHouse struct {
	name string
}

// 不能使用指针???
// 非侵入性接口实现,只需要实现方法,不需要实现接口
func (m Mouse) start()  {
	fmt.Println(m.name,"开始工作")
}
func (m Mouse) end()  {
	fmt.Println(m.name,"结束工作")
}
func (c ClickHouse) start()  {
	fmt.Println(c.name,"开始工作")
}
func (c ClickHouse) end()  {
	fmt.Println(c.name,"结束工作")
}

// 接口测试,多态
func testUSB(u USB)  {
	u.start()
	u.end()
}

func main() {
	m := Mouse{"鼠标"}
	c := ClickHouse{"键盘"}
	m.start()
	testUSB(m)
	testUSB(c)
}