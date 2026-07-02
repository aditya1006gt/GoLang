package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ /* ++i will not work in Go */ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}
	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println()
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println()
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}
}
