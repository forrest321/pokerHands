package pokerHands

import "fmt"

type Card struct {
	Rank Rank
	Suit Suit
	Rune rune
}

func (c *Card) String() string {
	return fmt.Sprintf("%s %s %v", c.Rank.Name, c.Suit.Symbol, c.Rune)
}

type Cards []Card

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Less(i, j int) bool {
	return c[i].Rank.Value < c[j].Rank.Value
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type Rank struct {
	Value   int
	Name    string
	Initial string
}

const (
	LowAceValue  = 1
	TwoValue     = 2
	ThreeValue   = 3
	FourValue    = 4
	FiveValue    = 5
	SixValue     = 6
	SevenValue   = 7
	EightValue   = 8
	NineValue    = 9
	TenValue     = 10
	JackValue    = 11
	QueenValue   = 12
	KingValue    = 13
	HighAceValue = 14
)

var (
	TwoRank     = Rank{TwoValue, "Two", "2"}
	ThreeRank   = Rank{ThreeValue, "Three", "3"}
	FourRank    = Rank{FourValue, "Four", "4"}
	FiveRank    = Rank{FiveValue, "Five", "5"}
	SixRank     = Rank{SixValue, "Six", "6"}
	SevenRank   = Rank{SevenValue, "Seven", "7"}
	EightRank   = Rank{EightValue, "Eight", "8"}
	NineRank    = Rank{NineValue, "Nine", "9"}
	TenRank     = Rank{TenValue, "Ten", "10"}
	JackRank    = Rank{JackValue, "Jack", "J"}
	QueenRank   = Rank{QueenValue, "Queen", "Q"}
	KingRank    = Rank{KingValue, "King", "K"}
	HighAceRank = Rank{HighAceValue, "Ace", "A"}
	LowAceRank  = Rank{LowAceValue, "Ace", "A"}
)

var Ranks = []Rank{
	//Low Ace is a special case, not included here to facilitate easy deck creation
	TwoRank,
	ThreeRank,
	FourRank,
	FiveRank,
	SixRank,
	SevenRank,
	EightRank,
	NineRank,
	TenRank,
	JackRank,
	QueenRank,
	KingRank,
	HighAceRank,
}

var (
	TwoOfClubs      = Card{Rank: TwoRank, Suit: Clubs, Rune: '\U0001F0D2'}
	ThreeOfClubs    = Card{Rank: ThreeRank, Suit: Clubs, Rune: '\U0001F0D3'}
	FourOfClubs     = Card{Rank: FourRank, Suit: Clubs, Rune: '\U0001F0D4'}
	FiveOfClubs     = Card{Rank: FiveRank, Suit: Clubs, Rune: '\U0001F0D5'}
	SixOfClubs      = Card{Rank: SixRank, Suit: Clubs, Rune: '\U0001F0D6'}
	SevenOfClubs    = Card{Rank: SevenRank, Suit: Clubs, Rune: '\U0001F0D7'}
	EightOfClubs    = Card{Rank: EightRank, Suit: Clubs, Rune: '\U0001F0D8'}
	NineOfClubs     = Card{Rank: NineRank, Suit: Clubs, Rune: '\U0001F0D9'}
	TenOfClubs      = Card{Rank: TenRank, Suit: Clubs, Rune: '\U0001F0DA'}
	JackOfClubs     = Card{Rank: JackRank, Suit: Clubs, Rune: '\U0001F0DB'}
	QueenOfClubs    = Card{Rank: QueenRank, Suit: Clubs, Rune: '\U0001F0DD'}
	KingOfClubs     = Card{Rank: KingRank, Suit: Clubs, Rune: '\U0001F0DE'}
	AceOfClubs      = Card{Rank: HighAceRank, Suit: Clubs, Rune: '\U0001F0D1'}
	TwoOfDiamonds   = Card{Rank: TwoRank, Suit: Diamonds, Rune: '\U0001F0C2'}
	ThreeOfDiamonds = Card{Rank: ThreeRank, Suit: Diamonds, Rune: '\U0001F0C3'}
	FourOfDiamonds  = Card{Rank: FourRank, Suit: Diamonds, Rune: '\U0001F0C4'}
	FiveOfDiamonds  = Card{Rank: FiveRank, Suit: Diamonds, Rune: '\U0001F0C5'}
	SixOfDiamonds   = Card{Rank: SixRank, Suit: Diamonds, Rune: '\U0001F0C6'}
	SevenOfDiamonds = Card{Rank: SevenRank, Suit: Diamonds, Rune: '\U0001F0C7'}
	EightOfDiamonds = Card{Rank: EightRank, Suit: Diamonds, Rune: '\U0001F0C8'}
	NineOfDiamonds  = Card{Rank: NineRank, Suit: Diamonds, Rune: '\U0001F0C9'}
	TenOfDiamonds   = Card{Rank: TenRank, Suit: Diamonds, Rune: '\U0001F0CA'}
	JackOfDiamonds  = Card{Rank: JackRank, Suit: Diamonds, Rune: '\U0001F0CB'}
	QueenOfDiamonds = Card{Rank: QueenRank, Suit: Diamonds, Rune: '\U0001F0CD'}
	KingOfDiamonds  = Card{Rank: KingRank, Suit: Diamonds, Rune: '\U0001F0CE'}
	AceOfDiamonds   = Card{Rank: HighAceRank, Suit: Diamonds, Rune: '\U0001F0C1'}
	TwoOfHearts     = Card{Rank: TwoRank, Suit: Hearts, Rune: '\U0001F0B2'}
	ThreeOfHearts   = Card{Rank: ThreeRank, Suit: Hearts, Rune: '\U0001F0B3'}
	FourOfHearts    = Card{Rank: FourRank, Suit: Hearts, Rune: '\U0001F0B4'}
	FiveOfHearts    = Card{Rank: FiveRank, Suit: Hearts, Rune: '\U0001F0B5'}
	SixOfHearts     = Card{Rank: SixRank, Suit: Hearts, Rune: '\U0001F0B6'}
	SevenOfHearts   = Card{Rank: SevenRank, Suit: Hearts, Rune: '\U0001F0B7'}
	EightOfHearts   = Card{Rank: EightRank, Suit: Hearts, Rune: '\U0001F0B8'}
	NineOfHearts    = Card{Rank: NineRank, Suit: Hearts, Rune: '\U0001F0B9'}
	TenOfHearts     = Card{Rank: TenRank, Suit: Hearts, Rune: '\U0001F0BA'}
	JackOfHearts    = Card{Rank: JackRank, Suit: Hearts, Rune: '\U0001F0BB'}
	QueenOfHearts   = Card{Rank: QueenRank, Suit: Hearts, Rune: '\U0001F0BD'}
	KingOfHearts    = Card{Rank: KingRank, Suit: Hearts, Rune: '\U0001F0BE'}
	AceOfHearts     = Card{Rank: HighAceRank, Suit: Hearts, Rune: '\U0001F0B1'}
	TwoOfSpades     = Card{Rank: TwoRank, Suit: Spades, Rune: '\U0001F0A2'}
	ThreeOfSpades   = Card{Rank: ThreeRank, Suit: Spades, Rune: '\U0001F0A3'}
	FourOfSpades    = Card{Rank: FourRank, Suit: Spades, Rune: '\U0001F0A4'}
	FiveOfSpades    = Card{Rank: FiveRank, Suit: Spades, Rune: '\U0001F0A5'}
	SixOfSpades     = Card{Rank: SixRank, Suit: Spades, Rune: '\U0001F0A6'}
	SevenOfSpades   = Card{Rank: SevenRank, Suit: Spades, Rune: '\U0001F0A7'}
	EightOfSpades   = Card{Rank: EightRank, Suit: Spades, Rune: '\U0001F0A8'}
	NineOfSpades    = Card{Rank: NineRank, Suit: Spades, Rune: '\U0001F0A9'}
	TenOfSpades     = Card{Rank: TenRank, Suit: Spades, Rune: '\U0001F0AA'}
	JackOfSpades    = Card{Rank: JackRank, Suit: Spades, Rune: '\U0001F0AB'}
	QueenOfSpades   = Card{Rank: QueenRank, Suit: Spades, Rune: '\U0001F0AD'}
	KingOfSpades    = Card{Rank: KingRank, Suit: Spades, Rune: '\U0001F0AE'}
	AceOfSpades     = Card{Rank: HighAceRank, Suit: Spades, Rune: '\U0001F0A1'}
)

var AllCards Cards = Cards{
	TwoOfClubs,
	ThreeOfClubs,
	FourOfClubs,
	FiveOfClubs,
	SixOfClubs,
	SevenOfClubs,
	EightOfClubs,
	NineOfClubs,
	TenOfClubs,
	JackOfClubs,
	QueenOfClubs,
	KingOfClubs,
	AceOfClubs,
	TwoOfDiamonds,
	ThreeOfDiamonds,
	FourOfDiamonds,
	FiveOfDiamonds,
	SixOfDiamonds,
	SevenOfDiamonds,
	EightOfDiamonds,
	NineOfDiamonds,
	TenOfDiamonds,
	JackOfDiamonds,
	QueenOfDiamonds,
	KingOfDiamonds,
	AceOfDiamonds,
	TwoOfHearts,
	ThreeOfHearts,
	FourOfHearts,
	FiveOfHearts,
	SixOfHearts,
	SevenOfHearts,
	EightOfHearts,
	NineOfHearts,
	TenOfHearts,
	JackOfHearts,
	QueenOfHearts,
	KingOfHearts,
	AceOfHearts,
	TwoOfSpades,
	ThreeOfSpades,
	FourOfSpades,
	FiveOfSpades,
	SixOfSpades,
	SevenOfSpades,
	EightOfSpades,
	NineOfSpades,
	TenOfSpades,
	JackOfSpades,
	QueenOfSpades,
	KingOfSpades,
	AceOfSpades,
}
