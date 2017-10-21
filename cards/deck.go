package cards

import (
	"math/rand"
)

// Deck represents an ordered list of Cards
type Deck struct {
	perm []int
	i    int
}

// NewShuffledDeck provides a Deck which has been shuffled
func NewShuffledDeck() Deck {
	return Deck{
		perm: rand.Perm(52),
	}
}

// Draw returns the Card at the top of the Deck, or an error if no cards remain
func (d *Deck) Draw() Cards {
	if d.i == len(d.perm) {
		return CardEmpty
	}

	d.i++
	return allCards[d.perm[d.i-1]]
}
