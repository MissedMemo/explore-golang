package main

import "fmt"

type colorMap map[string]string

func main() {

	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println(colors1)

	// nil map (can't be assigned to)
	//var colors2 map[string]string

	colors3 := make(map[string]string)
	colors3["gray1"] = "#7d7d7d"
	colors3["gray2"] = "#b4b4b4"

	fmt.Println(colors3)

	delete(colors3, "gray2")
	fmt.Println(colors3)

	colorsX := colorMap{
		"pink":   "#ff007f",
		"purple": "#663399",
	}

	printMap(colorsX)
	colorsX.print()
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("hex value for", color, "is", hex)
	}
}

func (m colorMap) print() {
	for color, hex := range m {
		fmt.Println("colorMap hue for", hex, "is", color)
	}
}
