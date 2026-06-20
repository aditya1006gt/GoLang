package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package in Golang")
	fmt.Println("Time package is used to work with date and time in Golang")

	presentTime := time.Now()
	fmt.Println("Present time is: ", presentTime)

	createdDate := time.Date(2004, time.June, 20, 8, 20, 35, 0, time.UTC)
	fmt.Println("Created date is: ", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
	//  GOOS="darwin" go build
}
