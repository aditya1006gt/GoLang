package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func main() {
	content := " Aditya, Creating a new file in Golang"
	path:= "./mygofile.txt"
	file, err := os.Create(path)
	// val, err := os.WriteFile("example.txt", []byte(content), 0644)
	// if err != nil {
	// 	fmt.Println("Error writing file:", err)
	// 	panic(err)
	// }
	CheckError(err)
	leng, err := io.WriteString(file, content)
	CheckError(err)
	fmt.Println("File written successfully:", leng)
	defer file.Close()

	readFile(path)
}

func readFile(path string) {
	databyte, err := ioutil.ReadFile(path)
	CheckError(err)
	fmt.Println("Raw data is:")
	fmt.Println(databyte)
	fmt.Println("File content:")
	fmt.Println(string(databyte))
}

// func readFile(path string) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		panic(err)
// 	}
// 	defer file.Close()

// 	content, err := io.ReadAll(file)
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		panic(err)
// 	}

// 	fmt.Println("File content:")
// 	fmt.Println(string(content))
// }
