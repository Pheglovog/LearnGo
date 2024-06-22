package main

import (
	"fmt"
)

// func main() {
// 	greeting := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
// 	fmt.Println(greeting)
// 	for _, v := range greeting {
// 		fmt.Println(v)
// 	}

// 	sub1 := greeting[:2]
// 	sub2 := greeting[2:4]
// 	sub3 := greeting[4:]
// 	fmt.Println(sub1)
// 	fmt.Println(sub2)
// 	fmt.Println(sub3)
// }

func main() {
	data := []string{"one", "two", "three"}
	var s1 []int
	s2 := make([]int, 5)
	s3 := make([]int, 0, 5)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	for _, v := range data {
		fmt.Println(v)
	}
}
