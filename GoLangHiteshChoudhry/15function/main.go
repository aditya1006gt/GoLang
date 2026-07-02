package main

import "fmt"

func add(x int, y int) int { // return type is int defined after the parameters
	return x + y
}

func ttladder(values ...int) int { // variadic function
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	fmt.Println("Welcome to Functions in Golang")
	greet()

	fmt.Println()

	addition := add(5, 5)
	fmt.Println("Addition is: ", addition)

	fmt.Println()

	total := ttladder(1, 2, 3, 4, 5)
	fmt.Println("Total is: ", total)

	fmt.Println()

	greeter()
}
func greeter() {
	fmt.Println("Hello from greeter function")
}
func greet() {
	fmt.Println("Hello from greet function")
}
