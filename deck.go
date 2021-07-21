package main

import "fmt"

// create a new type of deck which is a slice of strings
type deck []string // creates a new type of deck
// it is a slice that holds string objects

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() { // loops through cards, prints out value of each card
	for i, card := range d {
		fmt.Println(i, card)
	}

}
