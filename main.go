package main

func main() {
	// Declaring a map and initialise its key-value pairs inside
	_colors := colors{ //map[string]string {
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// // Declaring a map by using make function
	// colors := make(map[string]string)

	// // Set key-value pair on the map
	// colors["white"] = "#ffffff"

	// Get a value from a key-value pair from a map
	// fmt.Println("White color's hex is", _colors["white"])
	_colors.printValue("white")

	_colors.print()

	// Delete a key-value pair in a map
	delete(_colors, "white")

	_colors.print()
}
