package main

import (
	"os"
	"testing"
)

// what are the assertions that you're making for your program? the assumptions that you're making

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	if len(testDeck) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(testDeck))
	}
	if testDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as first element but got %v", testDeck[0])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := importDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
