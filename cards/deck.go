package main

import "fmt"

// Create a new type of deck
// which is a new slice of strings
type deck []string

func newDeck() deck {

	newDeck := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}

	return newDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
