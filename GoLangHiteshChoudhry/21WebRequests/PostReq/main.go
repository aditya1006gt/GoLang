package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Post req on GoLang url=http://localhost:3000/post")
	PerformPostRequest()
}

func PerformPostRequest() {
	const myurl = "http://localhost:3000/post"
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with Golang",
			"price": 0,
			"platform": "learncodeonline.in"
		}
	`)
	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
