// package main

// import (
// 	"fmt"
// )

// func main() {
// 	name := "Bill"
// 	fmt.Println(&name)
// }

// package main

// import "fmt"

// type location struct {
// 	longitude float64
// 	latitude  float64
// }

// func main() {
// 	newYork := location{
// 		latitude:  40.73,
// 		longitude: -73.93,
// 	}
// 	newYork.changeLatitude()
// 	fmt.Println(newYork)
// }

// func (lo *location) changeLatitude() {
// 	(*lo).latitude = 41.0
// }

package main

import "fmt"

func main() {
	name := "Bill"
	// fmt.Println(*&name)
	// fmt.Println(&name)
	fmt.Println(name)
}
