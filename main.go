package main

import "fmt"

func main() {
	// // Declaring a map and initialise its key-value pairs inside
	// colors := map[string]string {
	// 	"red":		"#ff0000",
	// 	"green":	"#4bf745",
	// }

	// Declaring a map by using make function
	colors := make(map[string]string)

	// Set key-value pair on the map
	colors["white"] = "#ffffff"

	// Get a value from a key-value pair from a map
	fmt.Println("White color's hex is", colors["white"])

	fmt.Println("Printed 'colors' map: ",colors)

	// Delete a key-value pair in a map
	delete(colors, "white")
		
	fmt.Println("Printed 'colors' map after deletion : ",colors)
}