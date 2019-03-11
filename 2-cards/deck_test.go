package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	numCardsExpected := 12

	if len(d) != numCardsExpected {
		t.Errorf("Expected %v cards in new deck, but got %v", numCardsExpected, len(d))
	}
}
