package main

import (
	"fmt"
	"math"
)

// 定于一个字符串
type areaError struct {
	msg string
	radius float64
}

func (a *areaError)Error() string{
	return fmt.Sprintf("error: 半径:%.2f,%s",a.radius,a.msg)
}

func calcArea(radius float64) (float64, error) {
	if radius<0{
		return 0,&areaError{"半径是非法的",radius}
	}
	return math.Pi*radius*radius, nil
}
func main() {
	//err := errors.New("新建")
	//fmt.Println(err)
	//fmt.Printf("%T\n",err)

	r := -1.0
	if area,err := calcArea(r); err==nil{
		fmt.Printf("面积:%.2f\n",area)
	}else {
		fmt.Println(err.Error())
	}
}
