package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	if a > b { // the { is mandatory after the condition
		// { if you add the { here after the condition, you will get an error
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}

	if num := 10; num%2 == 0 { // you can also declare a variable in the if statement
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}
}
