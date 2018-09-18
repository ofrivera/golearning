package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	//contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// fmt.Println(" ")
	//-----
	jim := person{
		firstName: "Jim",
		lastName:  "Part",
		contactInfo: contactInfo{
			email:   "jim@jim.com",
			zipCode: 94000,
		},
	}
	// fmt.Printf("%+v", jim)
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
