package cards

import (
	"testing"
)

func TestContains(t *testing.T) {
	var inout = map[struct {
		h Cards
		n Cards
	}]bool{
		{AceClubs, AceClubs}: true,
		{AceClubs, TwoClubs}: false,
		{TwoClubs, TwoClubs}: true,

		{AceClubs | TwoClubs, AceClubs}:   true,
		{AceClubs | TwoClubs, ThreeClubs}: false,

		{AceClubs | TwoClubs | ThreeClubs, AceClubs}:   true,
		{AceClubs | TwoClubs | ThreeClubs, TwoClubs}:   true,
		{AceClubs | TwoClubs | ThreeClubs, ThreeClubs}: true,

		{AceClubs | TwoClubs | ThreeClubs, TwoClubs | ThreeClubs}: true,

		{KingSpades, KingSpades}: true,
	}

	for in, out := range inout {
		if in.h.Contains(in.n) != out {
			t.Errorf("Contains(): %b, %b: %v", in.h, in.n, in.h.Contains(in.n))
		}
	}
}

func TestRank(t *testing.T) {
	var inout = map[Cards]Rank{
		AceClubs:    Ace,
		AceDiamonds: Ace,
		TwoDiamonds: Two,
		ThreeHearts: Three,
		KingSpades:  King,
	}

	for in, out := range inout {
		v := in.Rank()
		if v != out {
			t.Errorf("Rank(): %b, %b", v, out)

		}
	}

	defer func() {
		err := recover()
		if err == nil || err != ErrCheckOnEmptyCards {
			t.Errorf("Rank() should have panic'd with ErrCheckOnEmptyCards")
		}
	}()

	CardEmpty.Rank()
}

func TestCount(t *testing.T) {
	var inout = map[Cards]int{
		AceClubs:                                      1,
		AceClubs | TwoClubs:                           2,
		FourClubs | FiveClubs | SixClubs | SevenClubs: 4,
	}

	for in, out := range inout {
		if in.Count() != out {
			t.Errorf("Count(): %v, %v", in.Count(), out)
		}
	}
}

func TestSuit(t *testing.T) {
	var inout = map[Cards]Suit{
		AceClubs:   Clubs,
		TwoClubs:   Clubs,
		AceHearts:  Hearts,
		KingHearts: Hearts,
	}

	for in, out := range inout {
		v := in.Suit()

		if v != out {
			t.Errorf("Suit(): %b, %b", v, out)
		}
	}

	defer func() {
		err := recover()
		if err == nil || err != ErrCheckOnEmptyCards {
			t.Errorf("Suit() should have panic'd with ErrCheckOnEmptyCards")
		}
	}()

	CardEmpty.Suit()
}

func TestTake(t *testing.T) {
	var inout = map[Cards]Cards{
		AceClubs:                 AceClubs,
		AceClubs | TwoClubs:      AceClubs,
		TwoClubs | ThreeClubs:    ThreeClubs,
		QueenSpades | KingSpades: KingSpades,
		CardEmpty:                CardEmpty,
	}

	for in, out := range inout {
		v := in.Take()
		if v != out {
			t.Errorf("suit(): %b, %b", v, out)
		}
	}

	x := AceClubs
	x.Take()

	if x.Take() != CardEmpty {
		t.Errorf("take(): expected CardsEmpty")
	}
}

func TestAsSlice(t *testing.T) {
	var inout = map[Cards][]Cards{
		CardEmpty:                             []Cards{},
		AceClubs:                              []Cards{AceClubs},
		AceClubs | KingSpades:                 []Cards{AceClubs, KingSpades},
		AceClubs | KingSpades | EightDiamonds: []Cards{AceClubs, KingSpades, EightDiamonds},
		KingSpades:                            []Cards{KingSpades},
	}

	for in, out := range inout {
		v := in.asSlice()
		if !same(v, out) {
			t.Errorf("asSlice(): %v, %v", v, out)
		}
	}
}

func same(a, b []Cards) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestString(t *testing.T) {
	var inout = map[Cards]string{
		CardEmpty:                             "",
		AceClubs:                              "Ace_Clubs",
		AceClubs | KingSpades:                 "Ace_Clubs King_Spades",
		AceClubs | KingSpades | EightDiamonds: "Ace_Clubs King_Spades Eight_Diamonds",
	}

	for in, out := range inout {
		v := in.String()
		if v != out {
			t.Errorf("String(): %v, %v", v, out)
		}
	}
}
