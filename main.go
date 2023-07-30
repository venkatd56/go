package main

import "fmt"

func main() {
//	var card string = "Ace of Spades"  //Variable declaration and value assignement - long form syntax
	card := "Ace of Spades"			   // 2nd way of declaring and assigning value - compact form syntax
	card = "Five of diamonds"		// card variable already initialized in the above step, no need to use := again
	fmt.Println(card)
}


// variable declaration

// syntax: var card string = "value" 

// var : keyword for creating a variable
// card: variable name
// string: variable type, only string type can be assigned to the variable 

// Go is a statically typed language, whenever we define a variable with a type, and this type cannot be changed 

// types in go: bool, int, string, float64

// := only used when creating a new variable, if we are reassigning a new value to an existing variable, no need to use this notation