package main

import (
	"fmt"
)

type MyInt int

func main() {
	var a any
	var mine MyInt = 20
	a = mine
	i1 := int(a.(MyInt))
	fmt.Println(i1)
	i2 := a.(MyInt)
	fmt.Println(i2)
}
