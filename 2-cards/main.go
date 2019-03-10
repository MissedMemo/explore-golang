package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	/*
		hand, remainingCards := deal(cards, 5)

		hand.print()
		remainingCards.print()
	*/
}

func newCard() string {
	return "Five of Diamonds"
}
