package main

import "fmt"

type englishBot struct{}
type spanishhBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishhBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func printGreeting(sb spanishhBot) {
	fmt.Println(sb.getGreeting())
}

func (spanishhBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}