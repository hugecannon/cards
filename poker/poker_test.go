package poker

import (
	"testing"

	"github.com/hugecannon/cards/cards"
)

func TestRoyalFlush(t *testing.T) {
	var inout = map[cards.Cards]cards.Cards{
		cards.CardEmpty: cards.CardEmpty,
		cards.AceClubs:  cards.CardEmpty,
		cards.AceClubs | cards.TwoClubs | cards.ThreeClubs | cards.FourClubs | cards.FiveClubs: cards.CardEmpty,
		cards.TenClubs | cards.JackClubs | cards.QueenClubs | cards.KingClubs | cards.AceClubs: cards.TenClubs | cards.JackClubs | cards.QueenClubs | cards.KingClubs | cards.AceClubs,
	}

	for in, out := range inout {
		v := RoyalFlush(in)
		if v != out {
			t.Errorf("got %b, expected %b", v, out)
		}
	}
}
func TestStraightFlush(t *testing.T) {
	var inout = map[cards.Cards]cards.Cards{
		cards.CardEmpty: cards.CardEmpty,
		cards.AceClubs:  cards.CardEmpty,
		cards.AceClubs | cards.TwoClubs | cards.ThreeClubs | cards.FourClubs | cards.FiveClubs: cards.AceClubs | cards.TwoClubs | cards.ThreeClubs | cards.FourClubs | cards.FiveClubs,
		cards.TenClubs | cards.JackClubs | cards.QueenClubs | cards.KingClubs | cards.AceClubs: cards.TenClubs | cards.JackClubs | cards.QueenClubs | cards.KingClubs | cards.AceClubs,
	}

	for in, out := range inout {
		v := StraightFlush(in)
		if v != out {
			t.Errorf("got %b, expected %b", v, out)
		}
	}
}

func TestFlush(t *testing.T) {
	var inout = map[cards.Cards]cards.Cards{
		cards.CardEmpty: cards.CardEmpty,
		cards.TwoHearts | cards.SixDiamonds | cards.ThreeDiamonds | cards.TenDiamonds | cards.JackDiamonds:   cards.CardEmpty,
		cards.TwoHearts | cards.SixHearts | cards.KingHearts | cards.JackHearts | cards.SevenHearts:          cards.TwoHearts | cards.SixHearts | cards.KingHearts | cards.JackHearts | cards.SevenHearts,
		cards.AceDiamonds | cards.SixDiamonds | cards.FiveDiamonds | cards.KingDiamonds | cards.FourDiamonds: cards.AceDiamonds | cards.SixDiamonds | cards.FiveDiamonds | cards.KingDiamonds | cards.FourDiamonds,
	}

	for in, out := range inout {
		v := Flush(in)
		if v != out {
			t.Errorf("got %b, expected %b", v, out)
		}
	}
}
