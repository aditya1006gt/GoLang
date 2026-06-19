package main

import (
	"fmt"
	"time"
)

func main() {
	
	time := time.Now()
	fmt.Printf("Current time: %v\n", time)

	formattedTime := time.Format("02-01-2006")
	fmt.Println("Formatted time: ", formattedTime)
}