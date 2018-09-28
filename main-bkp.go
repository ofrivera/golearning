package main

import (
	"fmt"
)

func main() {
	//colors := make(map[string]string)
	//colors["white"] = "#fffff"

	// colors := make(map[int]string)
	// colors[10] = "#fffff"

	// 2
	//var colors map[string]string

	// 1
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf45",
	// }
	//fmt.Println(colors)

	// delete(colors, 10)
	// fmt.Println(colors)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf45",
	// 	"white": "#fffff",
	// }
	// printMap(colors)
	m := map[string]string{
		"dog": "bark",
	}
<<<<<<< HEAD
	changeMap(m)
	fmt.Println(m)
=======
	// fmt.Printf("%+v", jim)
	//jimPointer := &jim

	jim.updateName("Jimmy The Pos")
	jim.print()
>>>>>>> 17110f15ead928a60b783ae88d8cf1b4286f1ab5
}

func changeMap(m map[string]string) {
	m["cat"] = "purr"
}

func printMap(c map[string]string) {
	//key,value
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
