package main

import (
	"log"
)

func main() {
	a := 10
	b := 20
	sum := a + b
	log.Printf("The sum of %d and %d is %d", a, b, sum)

	c:= "adi"
	d:= "golang"
	concatenated := c + " " + d
	log.Printf("Concatenated string: %s", concatenated)

	var e string
	log.Printf("Value of e: %v", e)


	message := "All goroutines finished"
	log.Println(message)
}