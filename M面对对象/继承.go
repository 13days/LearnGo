package main

import "fmt"

// 父类
type Person struct {
	name string
	age int
}
// 子类
type Student struct {
	Person // 模拟继承
	school string
}

func (p *Person)talk()  {
	fmt.Println(p.name,"is talking.")
}

func main() {
	wcp := Person{name: "WCP", age: 23}
	p := Student{wcp,"GDUF"}
	fmt.Println(p.name,p.age,p.school)
	p = Student{Person{"wcp", 23}, "gduf"}
	fmt.Println(p)
	p.talk()
}
