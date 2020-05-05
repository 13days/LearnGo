package main

import "fmt"

func f1()  {

}

func f2(a int)  {

}

func f3(a,b,c int, x, y, z string) {

}
func f4(a,b int) int{
	return 0
}
func f5(a,b int) (int,int){
	return 0,0
}
func main() {
	fmt.Printf("%T\n",f1)
	fmt.Printf("%T\n",f2)
	fmt.Printf("%T\n",f3)
	fmt.Printf("%T\n",f4)
	fmt.Printf("%T\n",f5)
}
