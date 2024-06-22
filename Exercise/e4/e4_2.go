package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mySlice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, rand.Intn(100))
	}

	for _, v := range mySlice {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println(v, "Six")
		case v%2 == 0:
			fmt.Println(v, "Two")
		case v%3 == 0:
			fmt.Println(v, "Three")
		default:
			fmt.Println("Never mind")
		}
	}
}
