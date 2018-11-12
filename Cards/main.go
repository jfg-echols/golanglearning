package main

func main() {
	cards := importDeckFromFile("my_cards")
	cards.shuffleDeck()
	cards.printAllCards()

	// cards.printAllCards()

	// hand, remainingCards := dealHand(cards, 5)
	// hand.printAllCards()
	// println("hand^^remaining__")
	// remainingCards.printAllCards()
	// cards.saveToFile("my_cards")
	// cards.importDeckFromFile("my_cards")
}
