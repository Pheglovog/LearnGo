package main

import "fmt"

func main() {
	var message string
	message = "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println(string(runes[3]))
}
