package pokerHands

import (
	"testing"
)

// HighCardHand creates a high card hand
func HighCardHand() []Card {
	return []Card{
		{Rank: Ranks[0], Suit: Spades},   // 2 of Spades
		{Rank: Ranks[2], Suit: Hearts},   // 4 of Hearts
		{Rank: Ranks[4], Suit: Diamonds}, // 6 of Diamonds
		{Rank: Ranks[6], Suit: Clubs},    // 8 of Clubs
		{Rank: Ranks[10], Suit: Spades},  // Jack of Spades
	}
}

// OnePairHand creates a one pair hand
func OnePairHand() []Card {
	return []Card{
		{Rank: Ranks[2], Suit: Spades},   // 4 of Spades
		{Rank: Ranks[2], Suit: Hearts},   // 4 of Hearts
		{Rank: Ranks[4], Suit: Diamonds}, // 6 of Diamonds
		{Rank: Ranks[6], Suit: Clubs},    // 8 of Clubs
		{Rank: Ranks[10], Suit: Spades},  // Jack of Spades
	}
}

// TwoPairHand creates a two pair hand
func TwoPairHand() []Card {
	return []Card{
		{Rank: Ranks[2], Suit: Spades},   // 4 of Spades
		{Rank: Ranks[2], Suit: Hearts},   // 4 of Hearts
		{Rank: Ranks[6], Suit: Diamonds}, // 8 of Diamonds
		{Rank: Ranks[6], Suit: Clubs},    // 8 of Clubs
		{Rank: Ranks[10], Suit: Spades},  // Jack of Spades
	}
}

// ThreeOfAKindHand creates a three of a kind hand
func ThreeOfAKindHand() []Card {
	return []Card{
		{Rank: Ranks[2], Suit: Spades},   // 4 of Spades
		{Rank: Ranks[2], Suit: Hearts},   // 4 of Hearts
		{Rank: Ranks[2], Suit: Diamonds}, // 4 of Diamonds
		{Rank: Ranks[6], Suit: Clubs},    // 8 of Clubs
		{Rank: Ranks[10], Suit: Spades},  // Jack of Spades
	}
}

// StraightHand creates a straight hand
func StraightHand() []Card {
	return []Card{
		{Rank: Ranks[0], Suit: Spades},   // 2 of Spades
		{Rank: Ranks[1], Suit: Hearts},   // 3 of Hearts
		{Rank: Ranks[2], Suit: Diamonds}, // 4 of Diamonds
		{Rank: Ranks[3], Suit: Clubs},    // 5 of Clubs
		{Rank: Ranks[4], Suit: Spades},   // 6 of Spades
	}
}

func AceLowStraightHand() []Card {
	return []Card{
		{Rank: Ranks[12], Suit: Spades},  // Ace of Spades
		{Rank: Ranks[0], Suit: Spades},   // 2 of Spades
		{Rank: Ranks[1], Suit: Hearts},   // 3 of Hearts
		{Rank: Ranks[2], Suit: Diamonds}, // 4 of Diamonds
		{Rank: Ranks[3], Suit: Clubs},    // 5 of Clubs
	}
}

func AceHighStraightHand() []Card {
	return []Card{
		{Rank: Ranks[8], Suit: Diamonds}, // 10 of Diamonds
		{Rank: Ranks[9], Suit: Spades},   // Jack of Spades
		{Rank: Ranks[10], Suit: Spades},  // Queen of Spades
		{Rank: Ranks[11], Suit: Spades},  // King of Spades
		{Rank: Ranks[12], Suit: Spades},  // Ace of Spades
	}
}

// FlushHand creates a flush hand
func FlushHand() []Card {
	return []Card{
		{Rank: Ranks[0], Suit: Spades}, // 2 of Spades
		{Rank: Ranks[2], Suit: Spades}, // 4 of Spades
		{Rank: Ranks[4], Suit: Spades}, // 6 of Spades
		{Rank: Ranks[6], Suit: Spades}, // 8 of Spades
		{Rank: Ranks[8], Suit: Spades}, // 10 of Spades
	}
}

// FullHouseHand creates a full house hand
func FullHouseHand() []Card {
	return []Card{
		{Rank: Ranks[2], Suit: Spades},   // 4 of Spades
		{Rank: Ranks[2], Suit: Hearts},   // 4 of Hearts
		{Rank: Ranks[2], Suit: Diamonds}, // 4 of Diamonds
		{Rank: Ranks[6], Suit: Clubs},    // 8 of Clubs
		{Rank: Ranks[6], Suit: Spades},   // 8 of Spades
	}
}

// FourOfAKindHand creates a four of a kind hand
func FourOfAKindHand() []Card {
	return []Card{
		{Rank: Ranks[2], Suit: Spades},   // 4 of Spades
		{Rank: Ranks[2], Suit: Hearts},   // 4 of Hearts
		{Rank: Ranks[2], Suit: Diamonds}, // 4 of Diamonds
		{Rank: Ranks[2], Suit: Clubs},    // 4 of Clubs
		{Rank: Ranks[10], Suit: Spades},  // Jack of Spades
	}
}

// StraightFlushHand creates a straight flush hand
func StraightFlushHand() []Card {
	return []Card{
		{Rank: Ranks[0], Suit: Spades}, // 2 of Spades
		{Rank: Ranks[1], Suit: Spades}, // 3 of Spades
		{Rank: Ranks[2], Suit: Spades}, // 4 of Spades
		{Rank: Ranks[3], Suit: Spades}, // 5 of Spades
		{Rank: Ranks[4], Suit: Spades}, // 6 of Spades
	}
}

// RoyalFlushHand creates a royal flush hand
func RoyalFlushHand() []Card {
	return []Card{
		{Rank: Ranks[8], Suit: Spades},  // 10 of Spades
		{Rank: Ranks[9], Suit: Spades},  // Jack of Spades
		{Rank: Ranks[10], Suit: Spades}, // Queen of Spades
		{Rank: Ranks[11], Suit: Spades}, // King of Spades
		{Rank: Ranks[12], Suit: Spades}, // Ace of Spades
	}
}

func TestHand_Evaluate(t *testing.T) {
	type fields struct {
		Cards []Card
		Type  string
		Value HandRank
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "High Card", fields: fields{Cards: HighCardHand(), Type: "High Card", Value: HighCard}},
		{name: "One Pair", fields: fields{Cards: OnePairHand(), Type: "One Pair", Value: OnePair}},
		{name: "Two Pair", fields: fields{Cards: TwoPairHand(), Type: "Two Pair", Value: TwoPairs}},
		{name: "Three of a Kind", fields: fields{Cards: ThreeOfAKindHand(), Type: "Three of a Kind", Value: ThreeOfAKind}},
		{name: "Straight", fields: fields{Cards: StraightHand(), Type: "Straight", Value: Straight}},
		{name: "Low Ace Straight", fields: fields{Cards: AceLowStraightHand(), Type: "Straight", Value: Straight}},
		{name: "High Ace Straight", fields: fields{Cards: AceHighStraightHand(), Type: "Straight", Value: Straight}},
		{name: "Flush", fields: fields{Cards: FlushHand(), Type: "Flush", Value: Flush}},
		{name: "Full House", fields: fields{Cards: FullHouseHand(), Type: "Full House", Value: FullHouse}},
		{name: "Four of a Kind", fields: fields{Cards: FourOfAKindHand(), Type: "Four of a Kind", Value: FourOfAKind}},
		{name: "Straight Flush", fields: fields{Cards: StraightFlushHand(), Type: "Straight Flush", Value: StraightFlush}},
		{name: "Royal Flush", fields: fields{Cards: RoyalFlushHand(), Type: "Royal Flush", Value: RoyalFlush}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				Cards: tt.fields.Cards,
				Type:  tt.fields.Type,
				Value: tt.fields.Value,
			}
			h.Evaluate()
			if h.Type != tt.fields.Type {
				t.Errorf("Expected hand type %s, but got %s", tt.fields.Type, h.Type)
			}
			if h.Value != tt.fields.Value {
				t.Errorf("Expected hand value %d, but got %d", tt.fields.Value, h.Value)
			}
		})
	}
}

func Test_isFlush(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, false},
		{"Two Pair", args{TwoPairHand()}, false},
		{"Three of a Kind", args{ThreeOfAKindHand()}, false},
		{"Straight", args{StraightHand()}, false},
		{"Ace Low Straight", args{AceLowStraightHand()}, false},
		{"Ace High Straight", args{AceHighStraightHand()}, false},
		{"Flush", args{FlushHand()}, true},
		{"Full House", args{FullHouseHand()}, false},
		{"Four of a Kind", args{FourOfAKindHand()}, false},
		{"Straight Flush", args{StraightFlushHand()}, true},
		{"Royal Flush", args{RoyalFlushHand()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlush(tt.args.cards); got != tt.want {
				t.Errorf("isFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFourOfAKind(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, false},
		{"Two Pair", args{TwoPairHand()}, false},
		{"Three of a Kind", args{ThreeOfAKindHand()}, false},
		{"Straight", args{StraightHand()}, false},
		{"Ace Low Straight", args{AceLowStraightHand()}, false},
		{"Ace High Straight", args{AceHighStraightHand()}, false},
		{"Flush", args{FlushHand()}, false},
		{"Full House", args{FullHouseHand()}, false},
		{"Four of a Kind", args{FourOfAKindHand()}, true},
		{"Straight Flush", args{StraightFlushHand()}, false},
		{"Royal Flush", args{RoyalFlushHand()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//count := getCardRankCount(tt.args.cards)
			//if got := isFourOfAKind(count); got != tt.want {
			//	t.Errorf("isFourOfAKind() = %v, want %v", got, tt.want)
			//}
			if got := isFourOfAKind(tt.args.cards); got != tt.want {
				t.Errorf("isFourOfAKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFullHouse(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, false},
		{"Two Pair", args{TwoPairHand()}, false},
		{"Three of a Kind", args{ThreeOfAKindHand()}, false},
		{"Straight", args{StraightHand()}, false},
		{"Ace Low Straight", args{AceLowStraightHand()}, false},
		{"Ace High Straight", args{AceHighStraightHand()}, false},
		{"Flush", args{FlushHand()}, false},
		{"Full House", args{FullHouseHand()}, true},
		{"Four of a Kind", args{FourOfAKindHand()}, false},
		{"Straight Flush", args{StraightFlushHand()}, false},
		{"Royal Flush", args{RoyalFlushHand()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//count := getCardRankCount(tt.args.cards)
			//if got := isFullHouse(count); got != tt.want {
			//	t.Errorf("isFullHouse() = %v, want %v", got, tt.want)
			//}
			if got := isFullHouse(tt.args.cards); got != tt.want {
				t.Errorf("isFullHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOnePair(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, true},
		{"Two Pair", args{TwoPairHand()}, false},
		{"Three of a Kind", args{ThreeOfAKindHand()}, false},
		{"Straight", args{StraightHand()}, false},
		{"Ace Low Straight", args{AceLowStraightHand()}, false},
		{"Ace High Straight", args{AceHighStraightHand()}, false},
		{"Flush", args{FlushHand()}, false},
		{"Full House", args{FullHouseHand()}, false},
		{"Four of a Kind", args{FourOfAKindHand()}, false},
		{"Straight Flush", args{StraightFlushHand()}, false},
		{"Royal Flush", args{RoyalFlushHand()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOnePair(tt.args.cards); got != tt.want {
				t.Errorf("isOnePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRoyalFlush(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, false},
		{"Two Pair", args{TwoPairHand()}, false},
		{"Three of a Kind", args{ThreeOfAKindHand()}, false},
		{"Straight", args{StraightHand()}, false},
		{"Ace Low Straight", args{AceLowStraightHand()}, false},
		{"Ace High Straight", args{AceHighStraightHand()}, false},
		{"Flush", args{FlushHand()}, false},
		{"Full House", args{FullHouseHand()}, false},
		{"Four of a Kind", args{FourOfAKindHand()}, false},
		{"Straight Flush", args{StraightFlushHand()}, false},
		{"Royal Flush", args{RoyalFlushHand()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRoyalFlush(tt.args.cards); got != tt.want {
				t.Errorf("isRoyalFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStraight(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, false},
		{"Two Pair", args{TwoPairHand()}, false},
		{"Three of a Kind", args{ThreeOfAKindHand()}, false},
		{"Straight", args{StraightHand()}, true},
		{"Ace Low Straight", args{AceLowStraightHand()}, true},
		{"Ace High Straight", args{AceHighStraightHand()}, true},
		{"Flush", args{FlushHand()}, false},
		{"Full House", args{FullHouseHand()}, false},
		{"Four of a Kind", args{FourOfAKindHand()}, false},
		{"Straight Flush", args{StraightFlushHand()}, true},
		{"Royal Flush", args{RoyalFlushHand()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.cards); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStraightFlush(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, false},
		{"Two Pair", args{TwoPairHand()}, false},
		{"Three of a Kind", args{ThreeOfAKindHand()}, false},
		{"Straight", args{StraightHand()}, false},
		{"Ace Low Straight", args{AceLowStraightHand()}, false},
		{"Ace High Straight", args{AceHighStraightHand()}, false},
		{"Flush", args{FlushHand()}, false},
		{"Full House", args{FullHouseHand()}, false},
		{"Four of a Kind", args{FourOfAKindHand()}, false},
		{"Straight Flush", args{StraightFlushHand()}, true},
		{"Royal Flush", args{RoyalFlushHand()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraightFlush(tt.args.cards); got != tt.want {
				t.Errorf("isStraightFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isThreeOfAKind(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, false},
		{"Two Pair", args{TwoPairHand()}, false},
		{"Three of a Kind", args{ThreeOfAKindHand()}, true},
		{"Straight", args{StraightHand()}, false},
		{"Ace Low Straight", args{AceLowStraightHand()}, false},
		{"Ace High Straight", args{AceHighStraightHand()}, false},
		{"Flush", args{FlushHand()}, false},
		{"Full House", args{FullHouseHand()}, false},
		{"Four of a Kind", args{FourOfAKindHand()}, false},
		{"Straight Flush", args{StraightFlushHand()}, false},
		{"Royal Flush", args{RoyalFlushHand()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//count := getCardRankCount(tt.args.cards)
			//if got := isThreeOfAKind(count); got != tt.want {
			//	t.Errorf("isThreeOfAKind() = %v, want %v", got, tt.want)
			//}
			if got := isThreeOfAKind(tt.args.cards); got != tt.want {
				t.Errorf("isThreeOfAKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTwoPair(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"High Card", args{HighCardHand()}, false},
		{"One Pair", args{OnePairHand()}, false},
		{"Two Pair", args{TwoPairHand()}, true},
		{"Three of a Kind", args{ThreeOfAKindHand()}, false},
		{"Straight", args{StraightHand()}, false},
		{"Ace Low Straight", args{AceLowStraightHand()}, false},
		{"Ace High Straight", args{AceHighStraightHand()}, false},
		{"Flush", args{FlushHand()}, false},
		{"Full House", args{FullHouseHand()}, false},
		{"Four of a Kind", args{FourOfAKindHand()}, false},
		{"Straight Flush", args{StraightFlushHand()}, false},
		{"Royal Flush", args{RoyalFlushHand()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTwoPair(tt.args.cards); got != tt.want {
				t.Errorf("isTwoPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
