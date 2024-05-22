package pokerHands

import (
	"slices"
	"testing"
)

func HighCardHand() Hand {
	return Hand{
		Cards: []Card{
			{Rank: JackRank, Suit: Spades},
			{Rank: TwoRank, Suit: Spades},   // 2 of Spades
			{Rank: FourRank, Suit: Hearts},  // 4 of Hearts
			{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
			{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
		},
		Type: HighCard,
		Used: []Card{{Rank: JackRank, Suit: Spades}},
		Unused: []Card{
			{Rank: TwoRank, Suit: Spades},   // 2 of Spades
			{Rank: FourRank, Suit: Hearts},  // 4 of Hearts
			{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
			{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
		},
	}
}
func LowerHighCardHand1() Hand {
	return Hand{
		Cards: HighCardLower1Cards(),
		Type:  HighCard,
		Used:  []Card{{Rank: NineRank, Suit: Spades}}, // 9 of Spades
		Unused: []Card{
			{Rank: TwoRank, Suit: Spades},   // 2 of Spades
			{Rank: FourRank, Suit: Hearts},  // 4 of Hearts
			{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
			{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
		},
	}
}
func LowerHighCardHand2() Hand {
	return Hand{
		Cards: HighCardLower2Cards(),
		Type:  HighCard,
		Used:  []Card{{Rank: NineRank, Suit: Spades}}, // 9 of Spades
		Unused: []Card{
			{Rank: TwoRank, Suit: Spades},    // 2 of Spades
			{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
			{Rank: FiveRank, Suit: Diamonds}, // 5 of Diamonds
			{Rank: EightRank, Suit: Clubs},   // 8 of Clubs
		},
	}
}

// HighCardCards creates a high card hand
func HighCardCards() []Card {
	return []Card{
		{Rank: TwoRank, Suit: Spades},   // 2 of Spades
		{Rank: FourRank, Suit: Hearts},  // 4 of Hearts
		{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
		{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
		{Rank: JackRank, Suit: Spades},  // Jack of Spades
	}
}
func HighCardLower1Cards() []Card {
	return []Card{
		{Rank: NineRank, Suit: Spades},  // 9 of Spades
		{Rank: TwoRank, Suit: Spades},   // 2 of Spades
		{Rank: FourRank, Suit: Hearts},  // 4 of Hearts
		{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
		{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
	}
}
func HighCardLower2Cards() []Card {
	return []Card{
		{Rank: NineRank, Suit: Spades},   // 9 of Spades
		{Rank: TwoRank, Suit: Spades},    // 2 of Spades
		{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
		{Rank: FiveRank, Suit: Diamonds}, // 5 of Diamonds
		{Rank: EightRank, Suit: Clubs},   // 8 of Clubs
	}
}

func OnePairHand() Hand {
	return Hand{
		Cards: OnePairCards(),
		Type:  OnePair,
		Used: []Card{
			{Rank: FourRank, Suit: Spades}, // 4 of Spades
			{Rank: FourRank, Suit: Hearts}, // 4 of Hearts
		},
		Unused: []Card{
			{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
			{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
			{Rank: JackRank, Suit: Spades},  // Jack of Spades
		},
	}
}
func OnePairHigh1Hand() Hand {
	return Hand{
		Cards: OnePairHigh1Cards(),
		Type:  OnePair,
		Used: []Card{
			{Rank: FiveRank, Suit: Spades}, // 5 of Spades
			{Rank: FiveRank, Suit: Hearts}, // 5 of Hearts
		},
		Unused: []Card{
			{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
			{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
			{Rank: JackRank, Suit: Spades},  // Jack of Spades
		},
	}
}
func OnePairHigh2Hand() Hand {
	return Hand{
		Cards: OnePairHigh2Cards(),
		Type:  OnePair,
		Used: []Card{
			{Rank: FiveRank, Suit: Spades}, // 5 of Spades
			{Rank: FiveRank, Suit: Hearts}, // 5 of Hearts
		},
		Unused: []Card{
			{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
			{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
			{Rank: QueenRank, Suit: Spades}, // Queen of Spades
		},
	}
}

// OnePairHand creates a one pair hand
func OnePairCards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},  // 4 of Spades
		{Rank: FourRank, Suit: Hearts},  // 4 of Hearts
		{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
		{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
		{Rank: JackRank, Suit: Spades},  // Jack of Spades
	}
}

// OnePairHand creates a one pair hand
func OnePairHigh1Cards() []Card {
	return []Card{
		{Rank: FiveRank, Suit: Spades},  // 5 of Spades
		{Rank: FiveRank, Suit: Hearts},  // 5 of Hearts
		{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
		{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
		{Rank: JackRank, Suit: Spades},  // Jack of Spades
	}
}

// OnePairHand creates a one pair hand
func OnePairHigh2Cards() []Card {
	return []Card{
		{Rank: FiveRank, Suit: Spades},  // 5 of Spades
		{Rank: FiveRank, Suit: Hearts},  // 5 of Hearts
		{Rank: SixRank, Suit: Diamonds}, // 6 of Diamonds
		{Rank: EightRank, Suit: Clubs},  // 8 of Clubs
		{Rank: QueenRank, Suit: Spades}, // Queen of Spades
	}
}

func TwoPairHand() Hand {
	return Hand{
		Cards: TwoPairCards(),
		Type:  TwoPairs,
		Used: []Card{
			{Rank: FourRank, Suit: Spades},    // 4 of Spades
			{Rank: FourRank, Suit: Hearts},    // 4 of Hearts
			{Rank: EightRank, Suit: Diamonds}, // 8 of Diamonds
			{Rank: EightRank, Suit: Clubs},    // 8 of Clubs
		},
		Unused: []Card{
			{Rank: JackRank, Suit: Spades}, // Jack of Spades
		},
	}
}
func TwoPairHigh1Hand() Hand {
	return Hand{
		Cards: TwoPairHigh1Cards(),
		Type:  TwoPairs,
		Used: []Card{
			{Rank: FourRank, Suit: Spades},    // 4 of Spades
			{Rank: FourRank, Suit: Hearts},    // 4 of Hearts
			{Rank: EightRank, Suit: Diamonds}, // 8 of Diamonds
			{Rank: EightRank, Suit: Clubs},    // 8 of Clubs
		},
		Unused: []Card{
			{Rank: QueenRank, Suit: Spades}, // Queen of Spades
		},
	}
}
func TwoPairHigh2Hand() Hand {
	return Hand{
		Cards: TwoPairHigh2Cards(),
		Type:  TwoPairs,
		Used: []Card{
			{Rank: FiveRank, Suit: Spades},    // 5 of Spades
			{Rank: FiveRank, Suit: Hearts},    // 5 of Hearts
			{Rank: EightRank, Suit: Diamonds}, // 8 of Diamonds
			{Rank: EightRank, Suit: Clubs},    // 8 of Clubs
		},
		Unused: []Card{
			{Rank: JackRank, Suit: Spades}, // Jack of Spades
		},
	}
}

// TwoPairHand creates a two pair hand
func TwoPairCards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},    // 4 of Spades
		{Rank: FourRank, Suit: Hearts},    // 4 of Hearts
		{Rank: EightRank, Suit: Diamonds}, // 8 of Diamonds
		{Rank: EightRank, Suit: Clubs},    // 8 of Clubs
		{Rank: JackRank, Suit: Spades},    // Jack of Spades
	}
}

// TwoPairHand creates a two pair hand
func TwoPairHigh1Cards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},    // 4 of Spades
		{Rank: FourRank, Suit: Hearts},    // 4 of Hearts
		{Rank: EightRank, Suit: Diamonds}, // 8 of Diamonds
		{Rank: EightRank, Suit: Clubs},    // 8 of Clubs
		{Rank: QueenRank, Suit: Spades},   // Queen of Spades
	}
}

// TwoPairHand creates a two pair hand
func TwoPairHigh2Cards() []Card {
	return []Card{
		{Rank: FiveRank, Suit: Spades},    // 5 of Spades
		{Rank: FiveRank, Suit: Hearts},    // 5 of Hearts
		{Rank: EightRank, Suit: Diamonds}, // 8 of Diamonds
		{Rank: EightRank, Suit: Clubs},    // 8 of Clubs
		{Rank: JackRank, Suit: Spades},    // Jack of Spades
	}
}

func ThreeOfAKindHand() Hand {
	return Hand{
		Cards: ThreeOfAKindCards(),
		Type:  ThreeOfAKind,
		Used: []Card{
			{Rank: FourRank, Suit: Spades},   // 4 of Spades
			{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
			{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		},
		Unused: []Card{
			{Rank: EightRank, Suit: Clubs}, // 8 of Clubs
			{Rank: JackRank, Suit: Spades}, // Jack of Spades
		},
	}
}
func ThreeOfAKindHigh1Hand() Hand {
	return Hand{
		Cards: ThreeOfAKindHigh1Cards(),
		Type:  ThreeOfAKind,
		Used: []Card{
			{Rank: FourRank, Suit: Spades},   // 4 of Spades
			{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
			{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		},
		Unused: []Card{
			{Rank: NineRank, Suit: Clubs},  // 9 of Clubs
			{Rank: JackRank, Suit: Spades}, // Jack of Spades
		},
	}
}
func ThreeOfAKindHigh2Hand() Hand {
	return Hand{
		Cards: ThreeOfAKindHigh2Cards(),
		Type:  ThreeOfAKind,
		Used: []Card{
			{Rank: FiveRank, Suit: Spades},   // 5 of Spades
			{Rank: FiveRank, Suit: Hearts},   // 5 of Hearts
			{Rank: FiveRank, Suit: Diamonds}, // 5 of Diamonds
		},
		Unused: []Card{
			{Rank: NineRank, Suit: Clubs},  // 8 of Clubs
			{Rank: JackRank, Suit: Spades}, // Jack of Spades
		},
	}
}

// ThreeOfAKindCards creates a three of a kind hand
func ThreeOfAKindCards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},   // 4 of Spades
		{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: EightRank, Suit: Clubs},   // 8 of Clubs
		{Rank: JackRank, Suit: Spades},   // Jack of Spades
	}
}

// ThreeOfAKindCards creates a three of a kind hand
func ThreeOfAKindHigh1Cards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},   // 4 of Spades
		{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: NineRank, Suit: Clubs},    // 9 of Clubs
		{Rank: JackRank, Suit: Spades},   // Jack of Spades
	}
}

// ThreeOfAKindCards creates a three of a kind hand
func ThreeOfAKindHigh2Cards() []Card {
	return []Card{
		{Rank: FiveRank, Suit: Spades},   // 5 of Spades
		{Rank: FiveRank, Suit: Hearts},   // 5 of Hearts
		{Rank: FiveRank, Suit: Diamonds}, // 5 of Diamonds
		{Rank: NineRank, Suit: Clubs},    // 8 of Clubs
		{Rank: JackRank, Suit: Spades},   // Jack of Spades
	}
}

func StraightHand() Hand {
	return Hand{
		Cards:  StraightCards(),
		Type:   Straight,
		Used:   StraightCards(),
		Unused: []Card{},
	}
}
func StraightHighHand() Hand {
	return Hand{
		Cards:  StraightHighCards(),
		Type:   Straight,
		Used:   StraightHighCards(),
		Unused: []Card{},
	}
}
func AceLowStraightHand() Hand {
	return Hand{
		Cards:  AceLowStraightCards(),
		Type:   Straight,
		Used:   AceLowStraightCards(),
		Unused: []Card{},
	}
}
func AceHighStraightHand() Hand {
	return Hand{
		Cards:  AceHighStraightCards(),
		Type:   Straight,
		Used:   AceHighStraightCards(),
		Unused: []Card{},
	}
}

// StraightCards creates a straight hand
func StraightCards() []Card {
	return []Card{
		{Rank: TwoRank, Suit: Spades},    // 2 of Spades
		{Rank: ThreeRank, Suit: Hearts},  // 3 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: FiveRank, Suit: Clubs},    // 5 of Clubs
		{Rank: SixRank, Suit: Spades},    // 6 of Spades
	}
}

// StraightCards creates a straight hand
func StraightHighCards() []Card {
	return []Card{
		{Rank: ThreeRank, Suit: Hearts},  // 3 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: FiveRank, Suit: Clubs},    // 5 of Clubs
		{Rank: SixRank, Suit: Spades},    // 6 of Spades
		{Rank: SevenRank, Suit: Spades},  // 7 of Spades
	}
}
func AceLowStraightCards() []Card {
	return []Card{
		{Rank: LowAceRank, Suit: Spades}, // Ace of Spades
		{Rank: TwoRank, Suit: Spades},    // 2 of Spades
		{Rank: ThreeRank, Suit: Hearts},  // 3 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: FiveRank, Suit: Clubs},    // 5 of Clubs
	}
}
func AceHighStraightCards() []Card {
	return []Card{
		{Rank: TenRank, Suit: Diamonds},   // 10 of Diamonds
		{Rank: JackRank, Suit: Spades},    // Jack of Spades
		{Rank: QueenRank, Suit: Spades},   // Queen of Spades
		{Rank: KingRank, Suit: Spades},    // King of Spades
		{Rank: HighAceRank, Suit: Spades}, // Ace of Spades
	}
}

func FlushHand() Hand {
	return Hand{
		Cards:  FlushCards(),
		Type:   Flush,
		Used:   FlushCards(),
		Unused: []Card{},
	}
}
func FlushHighHand() Hand {
	return Hand{
		Cards:  FlushHighCards(),
		Type:   Flush,
		Used:   FlushHighCards(),
		Unused: []Card{},
	}
}

// FlushCards creates a flush hand
func FlushCards() []Card {
	return []Card{
		{Rank: TwoRank, Suit: Spades},   // 2 of Spades
		{Rank: FourRank, Suit: Spades},  // 4 of Spades
		{Rank: SixRank, Suit: Spades},   // 6 of Spades
		{Rank: EightRank, Suit: Spades}, // 8 of Spades
		{Rank: TenRank, Suit: Spades},   // 10 of Spades
	}
}

// FlushCards creates a flush hand
func FlushHighCards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},  // 4 of Spades
		{Rank: SixRank, Suit: Spades},   // 6 of Spades
		{Rank: EightRank, Suit: Spades}, // 8 of Spades
		{Rank: TenRank, Suit: Spades},   // 10 of Spades
		{Rank: JackRank, Suit: Spades},  // Jack of Spades
	}
}

func FullHouseHand() Hand {
	return Hand{
		Cards:  FullHouseCards(),
		Type:   FullHouse,
		Used:   FullHouseCards(),
		Unused: []Card{},
	}
}
func FullHouseHigh1Hand() Hand {
	return Hand{
		Cards:  FullHouseHigh1Cards(),
		Type:   FullHouse,
		Used:   FullHouseHigh1Cards(),
		Unused: []Card{},
	}
}
func FullHouseHigh2Hand() Hand {
	return Hand{
		Cards:  FullHouseHigh2Cards(),
		Type:   FullHouse,
		Used:   FullHouseHigh2Cards(),
		Unused: []Card{},
	}
}

// FullHouseCards creates a full house hand
func FullHouseCards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},   // 4 of Spades
		{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: EightRank, Suit: Clubs},   // 8 of Clubs
		{Rank: EightRank, Suit: Spades},  // 8 of Spades
	}
}

// FullHouseCards creates a full house hand
func FullHouseHigh1Cards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},   // 4 of Spades
		{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: NineRank, Suit: Clubs},    // 9 of Clubs
		{Rank: NineRank, Suit: Spades},   // 9 of Spades
	}
}

// FullHouseCards creates a full house hand
func FullHouseHigh2Cards() []Card {
	return []Card{
		{Rank: FiveRank, Suit: Spades},   // 5 of Spades
		{Rank: FiveRank, Suit: Hearts},   // 5 of Hearts
		{Rank: FiveRank, Suit: Diamonds}, // 5 of Diamonds
		{Rank: EightRank, Suit: Clubs},   // 8 of Clubs
		{Rank: EightRank, Suit: Spades},  // 8 of Spades
	}
}

func FourOfAKindHand() Hand {
	return Hand{
		Cards: FourOfAKindCards(),
		Type:  FourOfAKind,
		Used: []Card{
			{Rank: FourRank, Suit: Spades},   // 4 of Spades
			{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
			{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
			{Rank: FourRank, Suit: Clubs},    // 4 of Clubs
		},
		Unused: []Card{
			{Rank: JackRank, Suit: Spades}, // Jack of Spades
		},
	}
}
func FourOfAKindHigh1Hand() Hand {
	return Hand{
		Cards: FourOfAKindHigh1Cards(),
		Type:  FourOfAKind,
		Used: []Card{
			{Rank: FourRank, Suit: Spades},   // 4 of Spades
			{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
			{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
			{Rank: FourRank, Suit: Clubs},    // 4 of Clubs
		},
		Unused: []Card{
			{Rank: QueenRank, Suit: Spades}, // Queen of Spades
		},
	}
}
func FourOfAKindHigh2Hand() Hand {
	return Hand{
		Cards: FourOfAKindHigh2Cards(),
		Type:  FourOfAKind,
		Used: []Card{
			{Rank: FiveRank, Suit: Spades},   // 5 of Spades
			{Rank: FiveRank, Suit: Hearts},   // 5 of Hearts
			{Rank: FiveRank, Suit: Diamonds}, // 5 of Diamonds
			{Rank: FiveRank, Suit: Clubs},    // 5 of Clubs
		},
		Unused: []Card{
			{Rank: QueenRank, Suit: Spades}, // Queen of Spades
		},
	}
}

// FourOfAKindCards creates a four of a kind hand
func FourOfAKindCards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},   // 4 of Spades
		{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: FourRank, Suit: Clubs},    // 4 of Clubs
		{Rank: JackRank, Suit: Spades},   // Jack of Spades
	}
}

// FourOfAKindCards creates a four of a kind hand
func FourOfAKindHigh1Cards() []Card {
	return []Card{
		{Rank: FourRank, Suit: Spades},   // 4 of Spades
		{Rank: FourRank, Suit: Hearts},   // 4 of Hearts
		{Rank: FourRank, Suit: Diamonds}, // 4 of Diamonds
		{Rank: FourRank, Suit: Clubs},    // 4 of Clubs
		{Rank: QueenRank, Suit: Spades},  // Queen of Spades
	}
}

// FourOfAKindCards creates a four of a kind hand
func FourOfAKindHigh2Cards() []Card {
	return []Card{
		{Rank: FiveRank, Suit: Spades},   // 5 of Spades
		{Rank: FiveRank, Suit: Hearts},   // 5 of Hearts
		{Rank: FiveRank, Suit: Diamonds}, // 5 of Diamonds
		{Rank: FiveRank, Suit: Clubs},    // 5 of Clubs
		{Rank: QueenRank, Suit: Spades},  // Queen of Spades
	}
}

func StraightFlushHand() Hand {
	return Hand{
		Cards:  StraightFlushCards(),
		Type:   StraightFlush,
		Used:   StraightFlushCards(),
		Unused: []Card{},
	}
}
func HighStraightFlushHand() Hand {
	return Hand{
		Cards:  StraightFlushHighCards(),
		Type:   StraightFlush,
		Used:   StraightFlushHighCards(),
		Unused: []Card{},
	}
}

// StraightFlushCards creates a straight flush hand
func StraightFlushCards() []Card {
	return []Card{
		{Rank: TwoRank, Suit: Spades},   // 2 of Spades
		{Rank: ThreeRank, Suit: Spades}, // 3 of Spades
		{Rank: FourRank, Suit: Spades},  // 4 of Spades
		{Rank: FiveRank, Suit: Spades},  // 5 of Spades
		{Rank: SixRank, Suit: Spades},   // 6 of Spades
	}
}

// StraightFlushCards creates a straight flush hand
func StraightFlushHighCards() []Card {
	return []Card{
		{Rank: ThreeRank, Suit: Spades}, // 3 of Spades
		{Rank: FourRank, Suit: Spades},  // 4 of Spades
		{Rank: FiveRank, Suit: Spades},  // 5 of Spades
		{Rank: SixRank, Suit: Spades},   // 6 of Spades
		{Rank: SevenRank, Suit: Spades}, // 7 of Spades
	}
}

func RoyalFlushHand() Hand {
	return Hand{
		Cards:  RoyalFlushCards(),
		Type:   RoyalFlush,
		Used:   RoyalFlushCards(),
		Unused: []Card{},
	}
}

// RoyalFlushCards creates a royal flush hand
func RoyalFlushCards() []Card {
	return []Card{
		{Rank: TenRank, Suit: Spades},     // 10 of Spades
		{Rank: JackRank, Suit: Spades},    // Jack of Spades
		{Rank: QueenRank, Suit: Spades},   // Queen of Spades
		{Rank: KingRank, Suit: Spades},    // King of Spades
		{Rank: HighAceRank, Suit: Spades}, // Ace of Spades
	}
}

func TestHand_Evaluate(t *testing.T) {
	type fields struct {
		Cards []Card
		Type  HandType
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "High Card", fields: fields{Cards: HighCardCards(), Type: HighCard}},
		{name: "High Card", fields: fields{Cards: HighCardLower1Cards(), Type: HighCard}},
		{name: "High Card", fields: fields{Cards: HighCardLower2Cards(), Type: HighCard}},
		{name: "One Pair", fields: fields{Cards: OnePairCards(), Type: OnePair}},
		{name: "One Pair", fields: fields{Cards: OnePairHigh1Cards(), Type: OnePair}},
		{name: "One Pair", fields: fields{Cards: OnePairHigh2Cards(), Type: OnePair}},
		{name: "Two Pair", fields: fields{Cards: TwoPairCards(), Type: TwoPairs}},
		{name: "Two Pair", fields: fields{Cards: TwoPairHigh1Cards(), Type: TwoPairs}},
		{name: "Two Pair", fields: fields{Cards: TwoPairHigh2Cards(), Type: TwoPairs}},
		{name: "Three of a Kind", fields: fields{Cards: ThreeOfAKindCards(), Type: ThreeOfAKind}},
		{name: "Three of a Kind", fields: fields{Cards: ThreeOfAKindHigh1Cards(), Type: ThreeOfAKind}},
		{name: "Three of a Kind", fields: fields{Cards: ThreeOfAKindHigh2Cards(), Type: ThreeOfAKind}},
		{name: "Straight", fields: fields{Cards: StraightCards(), Type: Straight}},
		{name: "Straight", fields: fields{Cards: StraightHighCards(), Type: Straight}},
		{name: "Straight", fields: fields{Cards: AceLowStraightCards(), Type: Straight}},
		{name: "Straight", fields: fields{Cards: AceHighStraightCards(), Type: Straight}},
		{name: "Flush", fields: fields{Cards: FlushCards(), Type: Flush}},
		{name: "Flush", fields: fields{Cards: FlushHighCards(), Type: Flush}},
		{name: "Full House", fields: fields{Cards: FullHouseCards(), Type: FullHouse}},
		{name: "Full House", fields: fields{Cards: FullHouseHigh1Cards(), Type: FullHouse}},
		{name: "Full House", fields: fields{Cards: FullHouseHigh2Cards(), Type: FullHouse}},
		{name: "Four of a Kind", fields: fields{Cards: FourOfAKindCards(), Type: FourOfAKind}},
		{name: "Four of a Kind", fields: fields{Cards: FourOfAKindHigh1Cards(), Type: FourOfAKind}},
		{name: "Four of a Kind", fields: fields{Cards: FourOfAKindHigh2Cards(), Type: FourOfAKind}},
		{name: "Straight Flush", fields: fields{Cards: StraightFlushCards(), Type: StraightFlush}},
		{name: "Straight Flush", fields: fields{Cards: StraightFlushHighCards(), Type: StraightFlush}},
		{name: "Royal Flush", fields: fields{Cards: RoyalFlushCards(), Type: RoyalFlush}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				Cards: tt.fields.Cards,
				Type:  tt.fields.Type,
			}
			h.Evaluate()
			if h.Type != tt.fields.Type {
				t.Errorf("Expected hand type %s, but got %s", tt.fields.Type.Name, h.Type.Name)
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
		{"High Card", args{HighCardCards()}, false},
		{"One Pair", args{OnePairCards()}, false},
		{"Two Pair", args{TwoPairCards()}, false},
		{"Three of a Kind", args{ThreeOfAKindCards()}, false},
		{"Straight", args{StraightCards()}, false},
		{"Ace Low Straight", args{AceLowStraightCards()}, false},
		{"Ace High Straight", args{AceHighStraightCards()}, false},
		{"Flush", args{FlushCards()}, true},
		{"Flush", args{FlushHighCards()}, true},
		{"Full House", args{FullHouseCards()}, false},
		{"Four of a Kind", args{FourOfAKindCards()}, false},
		{"Straight Flush", args{StraightFlushCards()}, true},
		{"Royal Flush", args{RoyalFlushCards()}, true},
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
		{"High Card", args{HighCardCards()}, false},
		{"One Pair", args{OnePairCards()}, false},
		{"Two Pair", args{TwoPairCards()}, false},
		{"Three of a Kind", args{ThreeOfAKindCards()}, false},
		{"Straight", args{StraightCards()}, false},
		{"Ace Low Straight", args{AceLowStraightCards()}, false},
		{"Ace High Straight", args{AceHighStraightCards()}, false},
		{"Flush", args{FlushCards()}, false},
		{"Full House", args{FullHouseCards()}, false},
		{"Four of a Kind", args{FourOfAKindCards()}, true},
		{"Four of a Kind", args{FourOfAKindHigh1Cards()}, true},
		{"Four of a Kind", args{FourOfAKindHigh2Cards()}, true},
		{"Straight Flush", args{StraightFlushCards()}, false},
		{"Royal Flush", args{RoyalFlushCards()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, count, _ := getRankings(tt.args.cards)
			if got, _ := isFourOfAKind(count); got != tt.want {
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
		{"High Card", args{HighCardCards()}, false},
		{"One Pair", args{OnePairCards()}, false},
		{"Two Pair", args{TwoPairCards()}, false},
		{"Three of a Kind", args{ThreeOfAKindCards()}, false},
		{"Straight", args{StraightCards()}, false},
		{"Ace Low Straight", args{AceLowStraightCards()}, false},
		{"Ace High Straight", args{AceHighStraightCards()}, false},
		{"Flush", args{FlushCards()}, false},
		{"Full House", args{FullHouseCards()}, true},
		{"Full House", args{FullHouseHigh1Cards()}, true},
		{"Full House", args{FullHouseHigh2Cards()}, true},
		{"Four of a Kind", args{FourOfAKindCards()}, false},
		{"Straight Flush", args{StraightFlushCards()}, false},
		{"Royal Flush", args{RoyalFlushCards()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, count, _ := getRankings(tt.args.cards)
			if got := isFullHouse(count); got != tt.want {
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
		{"High Card", args{HighCardCards()}, false},
		{"One Pair", args{OnePairCards()}, true},
		{"One Pair", args{OnePairHigh1Cards()}, true},
		{"One Pair", args{OnePairHigh2Cards()}, true},
		{"Two Pair", args{TwoPairCards()}, false},
		{"Three of a Kind", args{ThreeOfAKindCards()}, false},
		{"Straight", args{StraightCards()}, false},
		{"Ace Low Straight", args{AceLowStraightCards()}, false},
		{"Ace High Straight", args{AceHighStraightCards()}, false},
		{"Flush", args{FlushCards()}, false},
		{"Full House", args{FullHouseCards()}, false},
		{"Four of a Kind", args{FourOfAKindCards()}, false},
		{"Straight Flush", args{StraightFlushCards()}, false},
		{"Royal Flush", args{RoyalFlushCards()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, count, _ := getRankings(tt.args.cards)
			if got, _ := isOnePair(count); got != tt.want {
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
		{"High Card", args{HighCardCards()}, false},
		{"One Pair", args{OnePairCards()}, false},
		{"Two Pair", args{TwoPairCards()}, false},
		{"Three of a Kind", args{ThreeOfAKindCards()}, false},
		{"Straight", args{StraightCards()}, false},
		{"Ace Low Straight", args{AceLowStraightCards()}, false},
		{"Ace High Straight", args{AceHighStraightCards()}, false},
		{"Flush", args{FlushCards()}, false},
		{"Full House", args{FullHouseCards()}, false},
		{"Four of a Kind", args{FourOfAKindCards()}, false},
		{"Straight Flush", args{StraightFlushCards()}, false},
		{"Royal Flush", args{RoyalFlushCards()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := getCardValues(tt.args.cards)
			flush := isFlush(tt.args.cards)
			straight := isStraight(getCardValues(tt.args.cards))
			if got := isRoyalFlush(vals, flush, straight); got != tt.want {
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
		{"High Card", args{HighCardCards()}, false},
		{"One Pair", args{OnePairCards()}, false},
		{"Two Pair", args{TwoPairCards()}, false},
		{"Three of a Kind", args{ThreeOfAKindCards()}, false},
		{"Straight", args{StraightCards()}, true},
		{"Straight", args{StraightHighCards()}, true},
		{"Ace Low Straight", args{AceLowStraightCards()}, true},
		{"Ace High Straight", args{AceHighStraightCards()}, true},
		{"Flush", args{FlushCards()}, false},
		{"Full House", args{FullHouseCards()}, false},
		{"Four of a Kind", args{FourOfAKindCards()}, false},
		{"Straight Flush", args{StraightFlushCards()}, true},
		{"Royal Flush", args{RoyalFlushCards()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := getCardValues(tt.args.cards)
			if got := isStraight(vals); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
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
		{"High Card", args{HighCardCards()}, false},
		{"One Pair", args{OnePairCards()}, false},
		{"Two Pair", args{TwoPairCards()}, false},
		{"Three of a Kind", args{ThreeOfAKindCards()}, true},
		{"Three of a Kind", args{ThreeOfAKindHigh1Cards()}, true},
		{"Three of a Kind", args{ThreeOfAKindHigh2Cards()}, true},
		{"Straight", args{StraightCards()}, false},
		{"Ace Low Straight", args{AceLowStraightCards()}, false},
		{"Ace High Straight", args{AceHighStraightCards()}, false},
		{"Flush", args{FlushCards()}, false},
		{"Full House", args{FullHouseCards()}, false},
		{"Four of a Kind", args{FourOfAKindCards()}, false},
		{"Straight Flush", args{StraightFlushCards()}, false},
		{"Royal Flush", args{RoyalFlushCards()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, count, _ := getRankings(tt.args.cards)
			if got, _ := isThreeOfAKind(count); got != tt.want {
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
		{"High Card", args{HighCardCards()}, false},
		{"One Pair", args{OnePairCards()}, false},
		{"Two Pair", args{TwoPairCards()}, true},
		{"Two Pair", args{TwoPairHigh1Cards()}, true},
		{"Two Pair", args{TwoPairHigh2Cards()}, true},
		{"Three of a Kind", args{ThreeOfAKindCards()}, false},
		{"Straight", args{StraightCards()}, false},
		{"Ace Low Straight", args{AceLowStraightCards()}, false},
		{"Ace High Straight", args{AceHighStraightCards()}, false},
		{"Flush", args{FlushCards()}, false},
		{"Full House", args{FullHouseCards()}, false},
		{"Four of a Kind", args{FourOfAKindCards()}, false},
		{"Straight Flush", args{StraightFlushCards()}, false},
		{"Royal Flush", args{RoyalFlushCards()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, count, _ := getRankings(tt.args.cards)
			if got, _, _ := isTwoPair(count); got != tt.want {
				t.Errorf("isTwoPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deck := NewDeck()
		deck.Shuffle()
		var hand Hand
		hand.Cards = deck.Deal(0, 5)
		hand.Evaluate()
	}
}

func Test_findWinners(t *testing.T) {
	type args struct {
		hands []Hand
	}
	tests := []struct {
		name string
		args args
		want []Hand
	}{
		{"1 winner - High Card", args{[]Hand{HighCardHand(), LowerHighCardHand1()}}, []Hand{HighCardHand()}},
		{"1 winner - High Card", args{[]Hand{HighCardHand(), LowerHighCardHand1(), LowerHighCardHand2()}}, []Hand{HighCardHand()}},
		{"1 winner - Lower High Card", args{[]Hand{LowerHighCardHand1(), LowerHighCardHand2()}}, []Hand{LowerHighCardHand1()}},
		{"1 Winner - One Pair", args{[]Hand{OnePairHand(), OnePairHigh1Hand(), OnePairHigh2Hand()}}, []Hand{OnePairHigh2Hand()}},
		{"1 Winner - Two Pair", args{[]Hand{TwoPairHand(), TwoPairHigh1Hand(), TwoPairHigh2Hand()}}, []Hand{TwoPairHigh2Hand()}},
		{"1 Winner - Three of a Kind", args{[]Hand{ThreeOfAKindHand(), ThreeOfAKindHigh1Hand(), ThreeOfAKindHigh2Hand()}}, []Hand{ThreeOfAKindHigh2Hand()}},
		{"1 Winner - Straight", args{[]Hand{StraightHand(), StraightHighHand()}}, []Hand{StraightHighHand()}},
		{"1 Winner - Flush", args{[]Hand{FlushHand(), FlushHighHand()}}, []Hand{FlushHighHand()}},
		{"1 Winner - Full House", args{[]Hand{FullHouseHand(), FullHouseHigh1Hand(), FullHouseHigh2Hand()}}, []Hand{FullHouseHigh2Hand()}},
		{"1 Winner - Four of a Kind", args{[]Hand{FourOfAKindHand(), FourOfAKindHigh1Hand(), FourOfAKindHigh2Hand()}}, []Hand{FourOfAKindHigh2Hand()}},
		{"1 winner - Straight Flush", args{[]Hand{StraightFlushHand(), HighStraightFlushHand()}}, []Hand{HighStraightFlushHand()}},
		{"1 winner - Royal Flush", args{[]Hand{HighCardHand(), RoyalFlushHand()}}, []Hand{RoyalFlushHand()}},
		{"2 winners - Royal Flush", args{[]Hand{RoyalFlushHand(), RoyalFlushHand()}}, []Hand{RoyalFlushHand(), RoyalFlushHand()}},
		{"1 winner - Royal Flush", args{[]Hand{StraightFlushHand(), RoyalFlushHand()}}, []Hand{RoyalFlushHand()}},
		{"1 winner - Royal Flush", args{[]Hand{StraightFlushHand(), RoyalFlushHand(), FourOfAKindHand(), FullHouseHigh2Hand(), StraightHand(), FlushHand(), ThreeOfAKindHigh1Hand(), TwoPairHigh2Hand(), OnePairHand(), HighCardHand()}}, []Hand{RoyalFlushHand()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindWinners(tt.args.hands)
			if len(got) == len(tt.want) {
				for i := 0; i < len(got); i++ {
					g, w := got[i], tt.want[i]
					if g.Type != w.Type {
						t.Errorf("Types - findWinners() = %v, want %v", &g, &w)
					}
					if !cardSlicesAreEqual(g.Cards, w.Cards) || !cardSlicesAreEqual(g.Used, w.Used) || !cardSlicesAreEqual(g.Unused, w.Unused) {
						t.Errorf("Slices - findWinners() = %v, want %v", &g, &w)
					}
				}
			} else {
				t.Errorf("findWinners() = %v, want %v", &got, &tt.want)
			}
		})
	}
}

func cardSlicesAreEqual(s1, s2 []Card) bool {
	if len(s1) != len(s2) {
		return false
	}
	v1, v2 := getCardValues(s1), getCardValues(s2)
	slices.Sort(v1)
	slices.Sort(v2)
	switch slices.Compare(v1, v2) {
	case 0:
		return true
	default:
		return false
	}
}
