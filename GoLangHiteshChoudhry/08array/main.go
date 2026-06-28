package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array") 
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	// fruitList[2] = "Grapes"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruit list is: ", len(fruitList))

	var vegList = [3]string{"Potato", "Tomato", "Brinjal"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Length of veg list is: ", len(vegList))
}