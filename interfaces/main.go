package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

//bad versions no interface
func (englishBot) getGreeting() string {
	return "Hello"
}
func (spanishBot) getGreeting() string {
	return "Hola"
}

//bad versions no interface
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	// makes it easier to reuse code
	//test program writes 2 chat bots - one in english and one in spanish
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}
