// Paul Schuster 3/4/17
// deck.go
// Models a deck of cards, with a value w/o context

package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Deck models a deck of cards all with a generic value
type Deck struct {
	numCards int
	cards    []int
}

// NewDeck is a factory for a deck
func NewDeck(num int) *Deck {
	d := new(Deck)
	d.numCards = num
	d.Shuffle()
	return d
}

// Resize and reorders the deck
func (d *Deck) Resize(x int) {
	d.numCards = x
	d.Shuffle()
}

// Check to see if one of the t cards are in the first num
func (d *Deck) Check(num, t int) bool {
	if num > len(d.cards) {
		num = len(d.cards)
	}
	for i := 0; i < num; i++ {
		if d.cards[i] < t {
			return true
		}
	}
	return false
}

// Shuffle reorders the deck randomly
func (d *Deck) Shuffle() {
	d.cards = rand.Perm(d.numCards)
}

// Ncheck returns how many till a target card from the begining of the deck
func (d *Deck) Ncheck(t int) int {
	for k, v := range d.cards {
		if v < t {
			return k
		}
	}
	return d.numCards
}
