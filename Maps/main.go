package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#12345",
	}

	colors["white"] = "#ffffff"

	printMap(colors)

	func printMap(c map)  {
		for color, hex := range c {
			fmt.Println("Hex code for", color, "is", hex)
		}
	}

}
