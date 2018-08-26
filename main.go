package main

func main() {
	// Creates a structure where [keys] are of type string, and values are of type string
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
	}

	//var emptyMap map[int]string

	otherColorMap := make(map[string]string)
	otherColorMap["white"] = "#FFFFFF"

	delete(colors, "red")

	PrintMap(colors)
	PrintMap(otherColorMap)
}

func PrintMap(c map[string]string) {
	for color, hex := range c {
		println("Key: " + color + " | Value: " + hex)
	}
}
