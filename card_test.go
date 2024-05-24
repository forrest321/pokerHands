package pokerHands

import "testing"

func TestCard_String(t *testing.T) {
	if len(AllCards) != 52 {
		t.Errorf("All cards should have 52 cards, but has %d", len(AllCards))
	}
}

func TestCard_Values(t *testing.T) {
	for _, card := range AllCards {
		if card.Rank.Value == 0 {
			t.Errorf("rank value should be non-zero, but has %d", card.Rank.Value)

		}
	}
}
