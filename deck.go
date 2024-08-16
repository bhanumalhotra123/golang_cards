package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	// func NewSource(seed int64) Source
	// NewSource returns a new pseudo-random Source seeded with the given value. Unlike the default Source used by top-level functions, this source is not safe for concurrent use by multiple goroutines. The returned Source implements Source64.
	source := rand.NewSource(time.Now().UnixNano())
	// func New(src Source) *Rand
	// New returns a new Rand that uses random values from src to generate other random values.
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1) // instead of rand packages version of intn we are going to use random number generator we created.

		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
