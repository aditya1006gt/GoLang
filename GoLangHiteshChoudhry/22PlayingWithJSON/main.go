package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // by doing this we can ignore the password field in the json output
	Tags     []string `json:"tags,omitempty"` // by doing this we can ignore the empty tags field in the json output
}

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	// EncodeJson()
	// fmt.Println("\n\n x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x \n\n")
	// EncodeJson2()
	fmt.Println("\n\n x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x \n\n")
	DecodeJson()
}

func EncodeJson() {
	adityaCourses := []Course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "learncodeonline.in", "abc123", nil},
		{"Angular Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
	}

	// package this data as JSON data
	finalJson, err := json.Marshal(adityaCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of %T and Value: %s\n", finalJson, finalJson)

	fmt.Println("---------------------------------------------------")

	finalJson1, err := json.MarshalIndent(adityaCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of %T and Value: %s\n", finalJson1, finalJson1)
}

func EncodeJson2() {
	adityaCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "learncodeonline.in", "abc123", nil},
		{"Angular Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
	}

	// package this data as JSON data
	finalJson, err := json.Marshal(adityaCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of %T and Value: %s\n", finalJson, finalJson)

	fmt.Println("---------------------------------------------------")

	finalJson1, err := json.MarshalIndent(adityaCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of %T and Value: %s\n", finalJson1, finalJson1)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "Angular Bootcamp",
                "Price": 299,
                "website": "learncodeonline.in",
                "tags": ["web-dev",
                        "js"
                ]
        }
	`)

	var adityaCourse Course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &adityaCourse)
		fmt.Printf("%#v\n", adityaCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key: %v and Value: %v and Type of Value: %T\n", k, v, v)
	}

}
