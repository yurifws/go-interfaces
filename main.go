package main

import "fmt"

type bot interface {
	getGreeting() string
}

type user struct{
	name string
}

type exampleInterface interface {
	getGreeting(string, int) (string, error)
	gotVersion() float64
	respondToUser(user) string
}

type englishBot struct{}
type spanishhBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishhBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishhBot) getGreeting() string {
	return "Hola!"
}