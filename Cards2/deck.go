package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck - slice of stringsd

type deck []string

func newFullDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardFaces := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, face := range cardFaces {
			cards = append(cards, face+" of "+suit)
		}
	}
	return cards
}

// d deck is a reciever - convention is a 1-3 letter abbreviation of the data type

func (d deck) printDeck() {
	for i, card := range d {

		fmt.Println(i, card)
	}
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// shuffledDeck := deck{}
	//for count deck get a random number from 0-deck length
	//get card from deck of that number
	//add it to shuffledDeck and remove it from deck

	//actually we're going through the deck and randomly swapping places with another
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
func dealHand(d deck, handsize int) (deck, deck) {
	// slice strings are object[startingindex:uptonotincluding]
	// in strings[] = {1,2,3,4}, strings[:2] would be {0,1}

	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveDeckoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newDeck := strings.Split(string(bs), ",")
	return deck(newDeck)
}
