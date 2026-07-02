package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:8000" // "https://minilambda.adixdevs.xyz/"

func main() {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	fmt.Printf("Type of Respnse %T\n", res)
	fmt.Println("Response ", res)
	// ⭐⭐⭐
	defer res.Body.Close() // Close the response body to free resources
	// very very Important to close the response body after reading it to avoid resource leaks.

	databyte, err := ioutil.ReadAll(res.Body) // Read the response body
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response data:", string(databyte))
}

// func main() {
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error fetching URL:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode != http.StatusOK {
// 		fmt.Println("Error: Status code", response.StatusCode)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	fmt.Println("Response:")
// 	fmt.Println(string(body))
// }
