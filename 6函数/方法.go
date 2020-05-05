package main

import "fmt"

/*定义结构体Circle*/
type Circle struct {
	radius float64
}
/*结构体Circle下的方法*/
func (c *Circle) getArea() float64{
	return 3.14*c.radius*c.radius
}
func main() {
	var c1 Circle
	c1.radius = 10
	fmt.Println(c1.getArea())
}
