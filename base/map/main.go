package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"balck": "#000000",
		"green": "#ffffff",
	}

	// add key vaule to map
	// colors["red"] = "#ff0000"
	// colors["white"] = "#ffffff"
	// delete key from map
	// delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Printf("key is %v, vaule is %v \n", k, v)
	}
}
