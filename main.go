package main

import "fmt"

func main() {
	card := newCard()

	card = "Five of Diamonds"

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
