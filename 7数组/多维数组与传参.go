package main

import "fmt"

func arrayTest()  {
	a := [][]int{
		{1,2,3},
		{3,4,5},
		{5,6,7},
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("%d ", a[i][j])
		}
		println()
	}
}

func getSum(a [][]float64) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			sum += a[i][j]
		}
	}
	return sum
}
func main() {
	//arrayTest()
	fmt.Println(getSum([][]float64{{1,2,3},{2,3,4},{1,23,3}}))
}
