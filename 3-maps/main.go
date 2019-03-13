package main

import "fmt"

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
}
