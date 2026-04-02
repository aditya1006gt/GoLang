package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  30,
	}
	var p2 Person
	p2.Name = "Bob"
	p2.Age = 25

	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Person 2: %+v\n", p2)
}