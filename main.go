package main

import (
	"fmt"
)

func main() {
	// Creates a structure where [keys] are of type string, and values are of type string
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
	}

	var emptyMap map[int]string

	otherColorMap := make(map[string]string)
	otherColorMap["white"] = "#FFFFFF"

	delete(colors, "red")

	fmt.Println(colors)
	fmt.Println(emptyMap)
	fmt.Println(otherColorMap)
}
