package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	worker1 := Employee{
		"Swift",
		"Talor",
		1,
	}
	worker2 := Employee{
		firstName: "Trump",
		lastName:  "Download",
		id:        2,
	}

	var worker3 Employee
	worker3.firstName = "Michal"
	worker3.lastName = "Nacy"
	worker3.id = 3

	fmt.Println(worker1)
	fmt.Println(worker2)
	fmt.Println(worker3)
}
