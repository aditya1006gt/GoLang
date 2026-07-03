package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Get req on GoLang url=http://localhost:3000/get")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"
	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	
	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Content length is: ", res.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)
	fmt.Println(content)

	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseString.String())
	fmt.Printf("type of response: %T\n", responseString.String())
	fmt.Println(string(content))
}



// byteCount, _ := responseString.Write(content)
	/*
	responseString.Write(content) copies the bytes in content into the strings.Builder.

	In your code:

	content is a byte slice returned by io.ReadAll.
	responseString is a buffer for building a string efficiently.
	Write appends those bytes to the builder.
	It returns how many bytes were written and an error.
	So this line is basically turning the response body bytes into the 
	builder’s internal string representation. Since you already have all 
	the bytes in content, this is a bit redundant unless you specifically 
	want to use strings.Builder. You could also just convert content 
	directly to a string when printing it.
	*/
