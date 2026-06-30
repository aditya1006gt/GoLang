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

	person2 := Person{
		Name:   "Bob",
		Age:    25,
		Gender: "Male",
	}
	fmt.Println("\nPerson 2:", person2)
	fmt.Printf("Person 2 details are: %+v\n", person2)
	fmt.Printf("Person 2 name is %v and Gender is %v", person2.Name, person2.Gender)
}

type Person struct {
	Name   string
	Age    int
	Gender string
}
