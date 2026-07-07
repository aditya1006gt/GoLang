package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	Platform string  `json:"platform"`
	Author   *Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var courses []Course

// {
// 	{ID: 1, Name: "ReactJS", Price: 299, Platform: "Udemy", Author: &Author{Name: "Hitesh Choudhry", Age: 30}},
// 	{ID: 2, Name: "Angular", Price: 199, Platform: "Udemy", Author: &Author{Name: "Hitesh Choudhry", Age: 30}},
// 	{ID: 3, Name: "VueJS", Price: 399, Platform: "Udemy", Author: &Author{Name: "Hitesh Choudhry", Age: 30}},
// 	{ID: 4, Name: "NodeJS", Price: 499, Platform: "Udemy", Author: &Author{Name: "Hitesh Choudhry", Age: 30}},
// }

func (c *Course) IsEmpty() bool {
	// return len(courses) == 0
	return c.ID == 0 && c.Name == ""
}

func main() {
	fmt.Println("Hello, start builing api's")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the HomePage!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	if len(courses) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No courses found"}`))
		return
	}

	json.NewEncoder(w).Encode(courses)
}
