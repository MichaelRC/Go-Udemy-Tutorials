package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	//Checks that first element is 'Ace of Spades'
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}
	//Checks that last element is 'King of Clubs'
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, but got %v", d[len(d)-1])
	}
}

//Long func name allows us to easily search for it
//using the func names it is trying to test
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//remove any old file named '_decktesting'
	os.Remove("_decktesting")

	//create a new deck, save it to file '_decktesting'
	deck := newDeck()
	deck.saveToFile("_decktesting")

	//load deck from file '_decktesting'
	loadedDeck := newDeckFromFile("_decktesting")

	//Make sure that the length of loaded deck is 52. If not, error.
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, got %v", len(loadedDeck))
	}
	//Remove '_decktesting' file for cleanup.
	os.Remove("_decktesting")
}
