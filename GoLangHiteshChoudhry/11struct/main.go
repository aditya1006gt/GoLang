package main

import "fmt"

func main() {
	person1 := Person{
		Name:   "Alice",
		Age:    30,
		Gender: "Female",
	}
	fmt.Println("Person 1:", person1)
	fmt.Printf("Person 1 details are: %+v\n", person1)
	fmt.Printf("Person 1 name is %v and Gender is %v", person1.Name, person1.Gender)
}

type Person struct {
	Name   string
	Age    int
	Gender string
}
