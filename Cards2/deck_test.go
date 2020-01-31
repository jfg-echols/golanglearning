package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newFullDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of spades first but got %v", d[0])
	}
	lastCard := d[len(d)-1]
	if lastCard != "King of Diamonds" {
		t.Errorf("Expected King of Diamonds but got %v", lastCard)
	}
	//function to test new deck
	//assume we should have 52 items in the new deck
	//assume the first card is ace of spades
	//assume last card is king of Diamonds
}

func TestSaveToDeckAndDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newFullDeck()
	deck.saveDeckoFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck to have 52 cards but got %v", len(loadedDeck))
	}
	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of spades first but got %v", loadedDeck[0])
	}
	lastCard := loadedDeck[len(loadedDeck)-1]
	if lastCard != "King of Diamonds" {
		t.Errorf("Expected King of Diamonds but got %v", lastCard)
	}

	os.Remove("_decktesting")
}
