package main

import "fmt"

type A interface {

}

func print(a ...A)  {
	for _, x := range a{
		fmt.Println(x)
	}
}

func main() {
	a := "1"
	b := 2
	c := true
	println(a,b,c)
}
