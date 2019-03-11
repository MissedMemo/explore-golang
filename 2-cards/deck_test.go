package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	numCardsExpected := 12

	if len(d) != numCardsExpected {
		t.Errorf("Expected %v cards in new deck, but got %v", numCardsExpected, len(d))
	}

	firstCard := d[0]
	expectedFirstCard := "one of diamonds"

	finalCard := d[len(d)-1]
	expectedFinalCard := "three of clubs"

	if firstCard != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but received %v", expectedFirstCard, firstCard)
	}

	if finalCard != expectedFinalCard {
		t.Errorf("Expected final card to be %v, but received %v", expectedFinalCard, finalCard)
	}
}
