package poker

import (
	"github.com/hugecannon/cards/cards"
)

func RoyalFlush(c cards.Cards) (ret cards.Cards) {
	ret = StraightFlush(c)

	var (
		hasAce  = ret&cards.Cards(cards.Ace) > 0
		hasKing = ret&cards.Cards(cards.King) > 0
	)

	if !(hasAce && hasKing) {
		ret = cards.CardEmpty
	}

	return
}

const sfMask uint64 = 0xF800000000000000

func StraightFlush(c cards.Cards) (ret cards.Cards) {
	// LowAce
	c |= ((c & cards.Cards(cards.Ace)) >> 13)
	for i := uint(0); i < 60; i++ {
		m := cards.Cards(sfMask >> i)
		if c&m == m {
			if m.Peek() > ret.Peek() {
				ret = m
			}
		}
	}

	// Low Aces back to High if we have LowAce straight
	lowAce := cards.Ace.C() >> 13
	if ret&cards.Two.C() > 0 {
		ret |= ((ret & lowAce) << 13)
		ret &= ^(lowAce)
	}

	return
}

func FourOfAKind(c cards.Cards) cards.Cards {
	for i := len(cards.Ranks) - 1; i >= 0; i-- {
		r := cards.Ranks[i].C()
		if c&r == r {
			return r
		}
	}

	return cards.CardEmpty
}

func FullHouse(c cards.Cards) cards.Cards {
	var (
		three = ThreeOfAKind(c)
		two   = Pair(c ^ three)
	)

	if three == cards.CardEmpty || two == cards.CardEmpty {
		return cards.CardEmpty
	}

	return three | two
}

// Flush returns the highest value 5-card flush in a Cards Set
func Flush(c cards.Cards) cards.Cards {
	highest := cards.CardEmpty
	for _, v := range cards.Suits {
		f := c & cards.Cards(v)

		if f.Count() >= 5 {
			if highest == cards.CardEmpty ||
				f.Peek() != cards.CardEmpty && f.Peek().Rank() > highest.Peek().Rank() {

				highest = cards.CardEmpty
				for i := 0; i < 5; i++ {
					highest |= f.Take()
				}
			}
		}
	}

	return highest
}

func Straight(c cards.Cards) cards.Cards {
	var (
		ret   cards.Cards
		count int
	)

	for i := len(cards.Ranks) - 1; i >= -1; i-- {
		var r cards.Cards
		if i == -1 {
			r = cards.Ace.C() // Low Ace
		} else {
			r = cards.Ranks[i].C()
		}

		d := c & r
		if d == cards.CardEmpty {
			count = 0
			ret = d
			continue
		}

		count++
		ret |= d.Take()

		if count == 5 {
			// Low Aces back to High
			ret |= ((ret & cards.Ace.C()) << 13)
			ret &= ^(cards.Ace.C() >> 13)

			return ret
		}
	}

	return cards.CardEmpty
}

func ThreeOfAKind(c cards.Cards) cards.Cards {
	for i := len(cards.Ranks) - 1; i >= 0; i-- {
		r := cards.Ranks[i].C()
		d := c & r

		if d.Count() == 3 {
			return d
		}
	}

	return cards.CardEmpty
}

func TwoPair(c cards.Cards) cards.Cards {
	var (
		p1 = Pair(c)
		p2 = Pair(c ^ p1)
	)

	if p1 == cards.CardEmpty || p2 == cards.CardEmpty {
		return cards.CardEmpty
	}

	return p1 | p2
}

func Pair(c cards.Cards) cards.Cards {
	for i := len(cards.Ranks) - 1; i >= 0; i-- {
		r := cards.Ranks[i].C()
		d := c & r

		if d.Count() == 2 {
			return d
		}
	}

	return cards.CardEmpty
}

func HighCard(c cards.Cards) cards.Cards {
	var ret cards.Cards

	for i := 0; i < 5; i++ {
		d := c.Take()
		if d == cards.CardEmpty {
			return cards.CardEmpty
		}
		ret |= d
	}

	return ret
}

func takeN(c cards.Cards, n int) (ret cards.Cards) {
	for i := 0; i < n; i++ {
		ret |= c.Take()
	}

	return
}

// BestHand works out the best 5-card hand for a set of cards
func BestHand(c cards.Cards) (Hand, cards.Cards) {
	if d := RoyalFlush(c); d != cards.CardEmpty {
		return HandRoyalFlush, d
	}

	if d := StraightFlush(c); d != cards.CardEmpty {
		return HandStraightFlush, d
	}

	if d := FourOfAKind(c); d != cards.CardEmpty {
		return HandFourOfAKind, d | (takeN(c^d, 1))
	}

	if d := FullHouse(c); d != cards.CardEmpty {
		return HandFullHouse, d
	}

	if d := Flush(c); d != cards.CardEmpty {
		return HandFlush, d
	}

	if d := Straight(c); d != cards.CardEmpty {
		return HandStraight, d
	}

	if d := ThreeOfAKind(c); d != cards.CardEmpty {
		return HandThreeOfAKind, d | (takeN(c^d, 2))
	}

	if d := TwoPair(c); d != cards.CardEmpty {
		return HandTwoPair, d | (takeN(c^d, 1))
	}

	if d := Pair(c); d != cards.CardEmpty {
		return HandPair, d | (takeN(c^d, 3))
	}

	return HandHighCard, HighCard(c)
}

// Hand represents one of the 10 hand types in Hold'Em poker
type Hand int64

func (h Hand) String() string {
	switch h {
	case HandHighCard:
		return "HighCard"
	case HandPair:
		return "Pair"
	case HandTwoPair:
		return "TwoPair"
	case HandThreeOfAKind:
		return "ThreeOfAKind"
	case HandStraight:
		return "Straight"
	case HandFlush:
		return "Flush"
	case HandFullHouse:
		return "FullHouse"
	case HandFourOfAKind:
		return "FourOfAKind"
	case HandStraightFlush:
		return "StraightFlush"
	case HandRoyalFlush:
		return "RoyalFlush"
	default:
		return "Invalid"
	}
}

const (
	HandHighCard Hand = iota
	HandPair
	HandTwoPair
	HandThreeOfAKind
	HandStraight
	HandFlush
	HandFullHouse
	HandFourOfAKind
	HandStraightFlush
	HandRoyalFlush
)
