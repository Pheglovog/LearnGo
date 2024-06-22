package main

import (
	"bytes"
	"fmt"
)

func prefixer(prefix string) func(string) string {
	return func(input string) string {
		var buffer bytes.Buffer
		buffer.WriteString(prefix)
		buffer.WriteString(" ")
		buffer.WriteString(input)
		return buffer.String()
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
