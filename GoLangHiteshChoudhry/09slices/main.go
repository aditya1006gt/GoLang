package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welecome to Slices\n\n")
	
	// Appending values to slices
	fmt.Println("Appending to Slices\n\n")

	var fruitList = []string{}
	fmt.Printf("Type of fruitList %T \n", fruitList)
	fruitList = append(fruitList, "Apple", "Orange", "Grapes", "Mango")
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruit list is: ", len(fruitList))

	fruitList = append(fruitList[1:])
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruit list is: ", len(fruitList))

	fruitList = append(fruitList[:2])
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruit list is: ", len(fruitList))


	hishScore := make([]int, 4)
	hishScore[0] = 234
	hishScore[1] = 345
	hishScore[2] = 456
	hishScore[3] = 567
	// highScore[4] = 678 // this will give error as we have defined length of slice as 4

	hishScore = append(hishScore, 678, 789, 890)
	fmt.Println("High Score is: ", hishScore)
	fmt.Println("Length of High Score is: ", len(hishScore))

	sort.Ints(hishScore)

	// Deleting a value from slice based on index

	fmt.Println("\n\nDeleting in Slices\n\n")

	var index int = 2
	hishScore = append(hishScore[:index], hishScore[index+1:]...)
	fmt.Println("High Score is: ", hishScore)
	fmt.Println("Length of High Score is: ", len(hishScore))
}
