package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Post req on GoLang url=http://localhost:3000/post")
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"
	// fake json payload

	reqBody:= strings.NewReader(`
		{
			"Aditya": "Let's go with Golang",
			"webDev":"courses",
			"pokemon":"games"
		}
	`)
	res, err := http.Post(myurl,"application/json",reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Aditya")
	data.Add("lastname", "Kumar")
	data.Add("email", "aditya.kumar@example.com")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}