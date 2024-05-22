package pokerHands

import "testing"

func TestCard_String(t *testing.T) {
	if len(AllCards) != 52 {
		t.Errorf("All cards should have 52 cards, but has %d", len(AllCards))
	}
}
