package cards

import "testing"

func TestDraw(t *testing.T) {
	d := NewShuffledDeck()
	for i := 0; i < 52; i++ {
		if d.Draw() == CardEmpty {
			t.Error("Didn't expect CardEmpty")
		}
	}

	if d.Draw() != CardEmpty {
		t.Error("Expected CardEmpty")
	}
}
