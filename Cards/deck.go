package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// craete a new type of 'deck' which is a slice of stinrgs ([]string)
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (deckVar deck) printAllCards() {
	for i, card := range deckVar {
		fmt.Println(i, card)
	}
}

func dealHand(deckVar deck, handSize int) (deck, deck) {
	return deckVar[:handSize], deckVar[handSize:]
}

func (deckVar deck) toString() string {
	return strings.Join([]string(deckVar), ",")

}
func (deckVar deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(deckVar.toString()), 0666)

}
func importDeckFromFile(fileName string) deck {
	importedByteSlice, importErr := ioutil.ReadFile(fileName)
	if importErr != nil {
		fmt.Println("Error: ", importErr)
		os.Exit(1)
	}
	s := strings.Split(string(importedByteSlice), ",")
	return deck(s)
}

func (deckVar deck) shuffleDeck() {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	for i := range deckVar {
		newPosition := random.Intn(len(deckVar) - 1)
		deckVar[i], deckVar[newPosition] = deckVar[newPosition], deckVar[i]
	}

}
