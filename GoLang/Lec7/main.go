package main

import (
	"fmt"
)

func main() {
	var age int= 30
	var name string= "Alice"
	var message= "All goroutines finished"
	const PI = 3.14159
	person := 123
	fmt.Println(message)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Value of PI: %.5f\n", PI)
	fmt.Printf("Person ID: %d\n", person)
}