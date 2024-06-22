package main

import (
	"fmt"
)

func UpdateSlice(arr []string, str string) {
	arr[len(arr)-1] = str
	fmt.Println(arr)
}

func GrowSlice(arr []string, str string) {
	arr = append(arr, str)
	fmt.Println(arr)
}

func main() {
	s := []string{"a", "b", "c"}
	UpdateSlice(s, "d")
	fmt.Println("in main after UpdateSlice:", s)
	GrowSlice(s, "e")
	fmt.Println("in main, after GrowSlice:", s)
}
