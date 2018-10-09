package main

import "fmt"

type bot interface {
	//Any type that includes a func called getGreeting and returns a type string are also honorary to bot type
	getGreeting() string
}

// type bot interface {
// 	getGreeting(string, int) (string,error)
// 	getBotVersion() float64
// 	respondToUser(user) string
// }

type englishBot struct{}
type spanishBot struct{}

func main() {
	//Honorary
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "HOoOla!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(eb spanishBot) {
// 	fmt.Println(eb.getGreeting())
// }
