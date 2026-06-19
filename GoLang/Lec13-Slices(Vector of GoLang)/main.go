package main

import (
	"log"
	"fmt"
)

func main() {
	
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	// Append a new number to the slice
	numbers = append(numbers, 6)
	fmt.Println("After appending 6:", numbers)

	// Create a new slice from the original slice
	newNumbers := numbers[1:4]
	fmt.Println("New slice (numbers[1:4]):", newNumbers)

	// Modify the new slice
	newNumbers[0] = 10
	fmt.Println("After modifying newNumbers[0]:", newNumbers)
	fmt.Println("Original slice after modification:", numbers)

	// Create a copy of the original slice
	copiedNumbers := make([]int, len(numbers))
	copy(copiedNumbers, numbers)
	fmt.Println("Copied slice:", copiedNumbers)

	// Modify the copied slice
	copiedNumbers[0] = 20
	fmt.Println("After modifying copiedNumbers[0]:", copiedNumbers)
	fmt.Println("Original slice after modifying copied slice:", numbers)

	// Log the final state of the original slice
	log.Printf("Final state of original slice: %v", numbers)
}