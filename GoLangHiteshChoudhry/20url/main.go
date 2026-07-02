package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://localhost:8000/learn/xxx?page=1&limit=10&sort=asc&filter=active"

func main() {
	fmt.Println("Welcome to handling URL's in Golang")
	fmt.Println(myurl)

	res, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Parsed URL:", res)

	fmt.Println("Scheme:", res.Scheme)
	fmt.Println("Host:", res.Host)
	fmt.Println("Path:", res.Path)
	fmt.Println("Raw Query:", res.RawQuery)

	queryParams := res.Query()
	fmt.Println("Query Parameters:", queryParams)
	fmt.Printf("Type of Query Parameters %T \n", queryParams)

	for key, values := range queryParams {
		// fmt.Printf("Query Parameter: %s = %s\n", key, values)
		// ^----- ye values slice hai, isliye hum loop lagayenge
		for _, value := range values {
			fmt.Printf("Query Parameter: %s = %s\n", key, value)
		}
	}
	fmt.Println();
	pathUrl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/path/to/resource",
		RawQuery: "param1=value1&param2=value2",
	} // *url.URL
	anotherurl := pathUrl.String() //string
	fmt.Println("Another URL:", anotherurl) 
	fmt.Println("Path URL:", pathUrl)
}
