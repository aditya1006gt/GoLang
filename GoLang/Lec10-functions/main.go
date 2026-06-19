package main

import (
	"log"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int, c float64) float64 {
	return float64(a - b) * c
}

func main() {
	log.Println("Hello, World!")
	result := add(3, 5)
	log.Printf("The result of adding 3 and 5 is: %d", result)
	subtractResult := subtract(10, 4, 2.5)
	log.Printf("The result of subtracting 4 from 10 and multiplying by 2.5 is: %.2f", subtractResult)
}