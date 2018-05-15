package main

// This code here is to demo the map concept within go similar to dict in python.

import (
	"fmt"
)

func main() {

	// One way to declare maps
	//colors := map[string]string{
	//	"red":  "0xfffff",
	//	"blue": "0xdecde0",
	//}

	//second way of declaring the maps
	//var colors map[string]string // initializes with the empty map.

	//3rd way of declaring the maps and populating it.
	colors := make(map[string]string)
	colors["white"] = "0x9034993"

	fmt.Println(colors)

	//Deleting the element from map
	delete(colors, "white")

	fmt.Println("After deletion :", colors)
}
