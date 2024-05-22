package pokerHands

import (
	"testing"
)

func TestDeck_Deal(t *testing.T) {
	d := NewDeck()
	d.Shuffle()
	tu := d.Deal(3, 4)
	if d.RemainingCards() != 45 {
		t.Errorf("remaining cards should be 45, got %d", d.RemainingCards())
	}
	if len(tu) != 4 {
		t.Errorf("expected 4 turn cards, got %d", len(tu))
	}
}

func TestDeck_DrawCard(t *testing.T) {
	d := NewDeck()
	d.Shuffle()
	c := d.DrawCard()
	if &c == nil {
		t.Error("expected non-nil card")
	}
	if c.Rank.Value < 2 {
		t.Errorf("expected rank value to be at least 2, got %d, card: %v", c.Rank.Value, c)
	}
	if d.RemainingCards() != 51 {
		t.Errorf("expected remaining cards to be 51, got %d", d.RemainingCards())
	}
}

func TestDeck_RemainingCards(t *testing.T) {
	d := NewDeck()
	if d.RemainingCards() != 52 {
		t.Errorf("remaining cards should be 52")
	}
	for i := range d.Cards {
		_ = d.DrawCard()
		if d.RemainingCards() != 51-i {
			t.Errorf("remaining cards should be %d", 51-i)
		}
	}
	if d.RemainingCards() != 0 {
		t.Errorf("remaining cards should be 0")
	}
}

func TestDeck_Shuffle(t *testing.T) {
	d := NewDeck()
	d.Shuffle()
	if len(d.Cards) != 52 {
		t.Errorf("Shuffle() did not preserve the number of cards")
	}
	if !d.Shuffled {
		t.Errorf("Shuffle() did not shuffle")
	}
}

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if d == nil {
		t.Errorf("NewDeck() returned nil")
	}
	if len(d.Cards) != 52 {
		t.Errorf("NewDeck() did not return 52 cards")
	}
	if d.Shuffled {
		t.Errorf("NewDeck() should not be shuffled")
	}
}
