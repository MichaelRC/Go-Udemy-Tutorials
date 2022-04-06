package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Create a new type of 'deck'
which is a slice of strings
*/
type deck []string

func newDeck() deck {
	//create a variable named cards
	//that is and empty deck
	cards := deck{}

	//crease a slice of strings
	//not a deck because they are not full cards
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//loop through Suits and Values and combine
	//each combination between them, adding it to cards deck
	//instead of useing i for index, used _ as it is an unused
	//variable. litterally a placeholder
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//allows us to add functionality to deck types
//to print them out without using fmt.Println
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
takes a deck type, and a int to determine 'handsize'
returns two deck(s), one up to handsize, and the other
is the rest.
*/
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

//turns deck into a slice of type string
//can do this because deck is a slice of strings
func (d deck) toString() string {
	/*
		convert deck (d) into a slice of type string
		then join the strings together seperated by a comma
		return the result
	*/
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	//bs = byte slice, err = error
	//collector variables for ReadFile returns
	bs, err := ioutil.ReadFile(filename)

	//error handeling if an error is returned
	//good practice to have this when error can be returned
	if err != nil {
		//print error returned along with string
		fmt.Println("!Error!: ", err)
		//passes an error code of 1, and exits program.
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	//random number generator
	//Calling the time package using the now function
	//and calling the UnixNano function on it.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	//we don't care about what we are
	//iterating over, just the index itself.
	for i := range d {
		//establish constraints based on the length
		//of the slice -1, pass it through the random
		//number generated in 'r' and store in newPosition
		newPosition := r.Intn(len(d) - 1)
		//swaps positions of i and newPosition in slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
