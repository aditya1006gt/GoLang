package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello, World!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)

	fmt.Printf("Type of name = %T ", name)

}
