package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating of the Pizza: ")
	rating, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for rating the Pizza %s", rating)

	numberRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	} else {
		fmt.Printf("Added 1 to your number rating = %f", numberRating+1)
	}
}
