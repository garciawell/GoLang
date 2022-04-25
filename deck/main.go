package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	cards.saveToFile("my_cards.csv")

	// cards2 := newDeckFromFile("my_cards")
	// cards2.print()

	cards3 := newDeck()
	cards3.shuffle()
	cards3.print()

}
