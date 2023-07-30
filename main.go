package main

import "fmt"

func main() {
//	var card string = "Ace of Spades"  //Variable declaration and value assignement - long form syntax
	card := newCard()			   // 2nd way of declaring and assigning value - compact form syntax
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}

// Functions
// These are similar to func main(), but we have to define the type of the value that the function returns
// if we are expexting a string, we need to pass the type string in the declaration as show in the above example

// func newCard() string {}