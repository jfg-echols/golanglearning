package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	//
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Ace of Hearts")
	cards := newFullDeck()
	cards.printDeck()
	// hand, remainingCards := dealHand(cards, 10)
	// hand.printDeck()
	// remainingCards.printDeck()
	cards.shuffleDeck()
	cards.printDeck()

	// cards.saveDeckoFile("my_cards")
	// newDeck := newDeckFromFile("my_cards")
	// newDeck.printDeck()

	// fmt.Println(cards)
}
