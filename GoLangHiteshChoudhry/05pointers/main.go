package main

import "fmt"

func main() {
	var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)
	// fmt.Printf("type of pointer is: %T\n", ptr)

	var a int = 10
	ptr = &a
	fmt.Println("Value of pointer is: ", ptr)
	fmt.Printf("type of pointer is: %T\n", ptr)

	fmt.Println("Value of a is: ", a, " and value of *ptr of a is: ", *ptr)
	fmt.Println("Value of a using pointer is: ", *ptr)

	*ptr = *ptr + 230
	fmt.Println("Value of a after incrementing using pointer is: ", a)

	fmt.Println("Value of a is: ", a, " and value of *ptr of a is: ", *ptr)
}
