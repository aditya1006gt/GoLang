package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Ludo Number Game\n ")
	var num int
	rand.Seed(time.Now().UnixNano())
	num = rand.Intn(6) + 1 // generates a random number between 1 and 6

	switch num {
	case 1:
		fmt.Println("You rolled a 1!")
	case 2:
		fmt.Println("You rolled a 2!")
	case 3:
		fmt.Println("You rolled a 3!")
	case 4:
		fmt.Println("You rolled a 4!")
		fallthrough // fallthrough is used to execute the next case even if the condition is not met
	case 5:
		fmt.Println("You rolled a 5!")
		fallthrough // fallthrough is used to execute the next case even if the condition is not met
	case 6:
		fmt.Println("You rolled a 6!")
		fallthrough
	default:
		fmt.Println("Invalid input! Please enter a number between 1 and 6.")
	}

}
