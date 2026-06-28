package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	// Creating a map
	fmt.Println("\n\n Creating a map \n\n")
	languages := make(map[string]string)

	// Adding key-value pairs to the map
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"

	fmt.Println("List of languages: ", languages)
	fmt.Println("Length of languages map: ", len(languages))

	// Deleting a key-value pair from the map

	fmt.Println("\n\n Deleting from a map \n\n")

	delete(languages, "RB")
	fmt.Println("List of languages after deleting Ruby: ", languages)
	fmt.Println("Length of languages map after deleting Ruby: ", len(languages))

	// Accessing a value from the map using a key

	fmt.Println("\n\n Accessing a value from the map using a key \n\n")

	value, isPresent := languages["JS"]
	if isPresent {
		fmt.Println("Value for key 'JS': ", value)
	} else {
		fmt.Println("Key 'JS' not found in the map")
	}
	fmt.Print("Value for key 'AA': ", languages["AA"]) // This will print the value for key "AA" if it exists, otherwise it will print nothing

	// Looping through a map

	fmt.Println("\n\nLooping through a map \n\n")

	for key, value := range languages {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
