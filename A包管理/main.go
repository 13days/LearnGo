package main // 包名可以和目录名不一样,但同个目录下的包名必须一样

import (
	t "byteDancer.com/zhongtai/GoLand_WuChuPeng/菜鸟教程/src/A包管理/lib" // 导入目录(模块)
	_ "byteDancer.com/zhongtai/LearnGo/A包管理/lib" // 调用init
	testx "byteDancer.com/zhongtai/LearnGo/A包管理/lib"
	"fmt"
	"myLib" // go/src下自定义的库
)

func main() {
	fmt.Println(myLib.Add(1,2))
	fmt.Println(t.Sub(5,6))
	testx.Nowa()
}
