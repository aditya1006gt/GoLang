package main

import "fmt"

// walrus operator not allowed outside function
// example: country := "Bharat" can't be declared outside function

const LoginToken string = "aditya123" // const is a global variable
// 1st letter of const should be capital if we want to use it outside the package
// it means it is public variable

func main() {
	var name string = "Aditya"
	var age uint8 = 25
	var isStudent bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)

	fmt.Printf("Variable is of type %T and value %v\n", name, name)
	fmt.Printf("Variable is of type %T and value %v\n", age, age)
	fmt.Printf("Variable is of type %T and value %v\n", isStudent, isStudent)

	// age = 266
	// fmt.Println("Updated Age:", age)

	// implicit type conversion
	var city = "Jharkhand"
	fmt.Println("City:", city)
	fmt.Printf("Variable is of type %T and value %v\n", city, city)

	// no var style
	// or
	// walrus operator
	country := "Bharat"
	fmt.Println("Country:", country)
	fmt.Printf("Variable is of type %T and value %v\n", country, country)

	fmt.Print("Login Token:", LoginToken)
}