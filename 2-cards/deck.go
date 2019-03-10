package main

import "fmt"

// define deck as a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"diamonds", "spades", "hearts", "clubs"}
	cardValues := []string{"one", "two", "three"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, numCards int) (deck, deck) {
	return d[:numCards], d[numCards:]
}
