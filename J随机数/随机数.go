package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int31n(100))
	}
}