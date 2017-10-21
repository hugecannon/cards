// Package cards provides concepts of Cards, Suits, and Ranks, along with some helpers for checking values
package cards

import (
	"fmt"
	"strings"
)

// Suit is a card suit
type Suit uint64

func (s Suit) C() Cards {
	return Cards(s)
}

// Rank is a card rank
type Rank uint64

func (r Rank) C() Cards {
	return Cards(r)
}

// Cards is a set of cards
type Cards uint64

// ErrCheckOnEmptyCards is caused when Suit() or Rank() is called on an Empty Cards set
var ErrCheckOnEmptyCards = fmt.Errorf("method called on CardEmpty Cards set")

// Contains checks a Cards set to see if it contains a Cards set
func (c Cards) Contains(ca Cards) bool {
	return (c & ca) == ca
}

func (c Cards) String() string {
	if c.Count() == 1 {
		r := c.Rank()
		s := c.Suit()
		return fmt.Sprintf("%s_%s", r, s)
	}

	return strings.Trim(fmt.Sprint(c.asSlice()), "[]") // :(
}

func (c Cards) asSlice() []Cards {
	var (
		ret = make([]Cards, c.Count())
		i   = 0
	)

	ca := c.Take()
	for ca != CardEmpty {
		ret[i] = ca

		ca = c.Take()
		i++
	}

	return ret
}

// Take removes and returns single Cards from a Cards set
func (c *Cards) Take() Cards {
	d := c.Peek()
	*c ^= d
	return d
}

// Peek returns the 'highest' value card from a Cards set
func (c Cards) Peek() Cards {
	for i := len(Ranks) - 1; i >= 0; i-- {
		r := Ranks[i].C()
		if d := c & r; d > 0 {
			return d.highest()
		}
	}

	return CardEmpty
}

func (c Cards) highest() Cards {
	c |= (c >> 1)
	c |= (c >> 2)
	c |= (c >> 4)
	c |= (c >> 8)
	c |= (c >> 16)
	c |= (c >> 32)

	return (c - (c >> 1))
}

// Suit returns the Suit of a single-carded Cards set
//
// Suit panics if invoked on a CardsEmpty cards set
func (c Cards) Suit() Suit {
	for _, v := range Suits {
		if c&Cards(v) > 0 {
			return v
		}
	}

	panic(ErrCheckOnEmptyCards)
}

// Rank returns the Rank of a single-carded Cards set
//
// Rank panics if invoked on a CardsEmpty cards set
func (c Cards) Rank() Rank {
	for _, v := range Ranks {
		if c&Cards(v) > 0 {
			return v
		}
	}

	panic(ErrCheckOnEmptyCards)
}

// Count returns the number of cards in a Cards set
func (c Cards) Count() (r int) {
	for c != 0 {
		c &= (c - 1)
		r++
	}

	return
}

func (s Suit) String() string {
	return suitStrings[s]
}

func (r Rank) String() string {
	return rankStrings[r]
}

const suitWidth = 16
const suitMask = ((1 << suitWidth) - 1)

const (
	Clubs    Suit = 0xFFFF
	Diamonds Suit = 0xFFFF0000
	Hearts   Suit = 0xFFFF00000000
	Spades   Suit = 0xFFFF000000000000

	clubsName    = "Clubs"
	diamondsName = "Diamonds"
	heartsName   = "Hearts"
	spadesName   = "Spades"
)

var (
	Suits = []Suit{
		Clubs,
		Diamonds,
		Hearts,
		Spades,
	}

	suitStrings = map[Suit]string{
		Clubs:    clubsName,
		Diamonds: diamondsName,
		Hearts:   heartsName,
		Spades:   spadesName,
	}
)

const (
	_ Cards = 1 << iota
	TwoClubs
	ThreeClubs
	FourClubs
	FiveClubs
	SixClubs
	SevenClubs
	EightClubs
	NineClubs
	TenClubs
	JackClubs
	QueenClubs
	KingClubs
	AceClubs
	_
	_
	_
	TwoDiamonds
	ThreeDiamonds
	FourDiamonds
	FiveDiamonds
	SixDiamonds
	SevenDiamonds
	EightDiamonds
	NineDiamonds
	TenDiamonds
	JackDiamonds
	QueenDiamonds
	KingDiamonds
	AceDiamonds
	_
	_
	_
	TwoHearts
	ThreeHearts
	FourHearts
	FiveHearts
	SixHearts
	SevenHearts
	EightHearts
	NineHearts
	TenHearts
	JackHearts
	QueenHearts
	KingHearts
	AceHearts
	_
	_
	_
	TwoSpades
	ThreeSpades
	FourSpades
	FiveSpades
	SixSpades
	SevenSpades
	EightSpades
	NineSpades
	TenSpades
	JackSpades
	QueenSpades
	KingSpades
	AceSpades

	CardEmpty = Cards(0)
)

const (
	Two   Rank = Rank(TwoClubs | TwoDiamonds | TwoHearts | TwoSpades)
	Three Rank = Rank(ThreeClubs | ThreeDiamonds | ThreeHearts | ThreeSpades)
	Four  Rank = Rank(FourClubs | FourDiamonds | FourHearts | FourSpades)
	Five  Rank = Rank(FiveClubs | FiveDiamonds | FiveHearts | FiveSpades)
	Six   Rank = Rank(SixClubs | SixDiamonds | SixHearts | SixSpades)
	Seven Rank = Rank(SevenClubs | SevenDiamonds | SevenHearts | SevenSpades)
	Eight Rank = Rank(EightClubs | EightDiamonds | EightHearts | EightSpades)
	Nine  Rank = Rank(NineClubs | NineDiamonds | NineHearts | NineSpades)
	Ten   Rank = Rank(TenClubs | TenDiamonds | TenHearts | TenSpades)
	Jack  Rank = Rank(JackClubs | JackDiamonds | JackHearts | JackSpades)
	Queen Rank = Rank(QueenClubs | QueenDiamonds | QueenHearts | QueenSpades)
	King  Rank = Rank(KingClubs | KingDiamonds | KingHearts | KingSpades)
	Ace   Rank = Rank(AceClubs | AceDiamonds | AceHearts | AceSpades)

	twoName   = "Two"
	threeName = "Three"
	fourName  = "Four"
	fiveName  = "Five"
	sixName   = "Six"
	sevenName = "Seven"
	eightName = "Eight"
	nineName  = "Nine"
	tenName   = "Ten"
	jackName  = "Jack"
	queenName = "Queen"
	kingName  = "King"
	aceName   = "Ace"
)

var (
	Ranks = []Rank{
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Ten,
		Jack,
		Queen,
		King,
		Ace,
	}

	rankStrings = map[Rank]string{
		Two:   twoName,
		Three: threeName,
		Four:  fourName,
		Five:  fiveName,
		Six:   sixName,
		Seven: sevenName,
		Eight: eightName,
		Nine:  nineName,
		Ten:   tenName,
		Jack:  jackName,
		Queen: queenName,
		King:  kingName,
		Ace:   aceName,
	}
)

var (
	allCards = []Cards{
		TwoClubs,
		ThreeClubs,
		FourClubs,
		FiveClubs,
		SixClubs,
		SevenClubs,
		EightClubs,
		NineClubs,
		TenClubs,
		JackClubs,
		QueenClubs,
		KingClubs,
		AceClubs,

		TwoDiamonds,
		ThreeDiamonds,
		FourDiamonds,
		FiveDiamonds,
		SixDiamonds,
		SevenDiamonds,
		EightDiamonds,
		NineDiamonds,
		TenDiamonds,
		JackDiamonds,
		QueenDiamonds,
		KingDiamonds,
		AceDiamonds,

		TwoHearts,
		ThreeHearts,
		FourHearts,
		FiveHearts,
		SixHearts,
		SevenHearts,
		EightHearts,
		NineHearts,
		TenHearts,
		JackHearts,
		QueenHearts,
		KingHearts,
		AceHearts,

		TwoSpades,
		ThreeSpades,
		FourSpades,
		FiveSpades,
		SixSpades,
		SevenSpades,
		EightSpades,
		NineSpades,
		TenSpades,
		JackSpades,
		QueenSpades,
		KingSpades,
		AceSpades,
	}
)
