package main

import "fmt"

type bot interface {
	getFreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getFreeting())
}

func (englishBot) getFreeting() string {
	// VERY CUSTOM LOGIC

	return "Hi There!"
}

func (spanishBot) getFreeting() string {
	return "Hola"
}
