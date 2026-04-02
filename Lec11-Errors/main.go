package main

import (
	"log"
	"fmt"
)

func subtract(a, b int, c float64) (float64, error) {
	if c == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	return float64(a - b) / c, nil
}

func main() {
	log.Println("Hello, World!")
	var a, b int
	var c float64

	fmt.Println("Enter two integers:")
	fmt.Scan(&a,&b);
	fmt.Println("Enter a float:")
	fmt.Scan(&c);

	subtractResult, err := subtract(a, b, c)
	if err != nil {
		log.Println("0")
		log.Fatal(err)
	}

	log.Println("1")
	log.Printf("The result of subtracting %d from %d and dividing by %.1f is: %.2f", b, a, c, subtractResult)

	log.Println("2")
	subdiv, _ := subtract(a, b, c)
	log.Printf("The result of subtracting %d from %d and dividing by %.1f is: %.2f", b, a, c, subdiv)
}