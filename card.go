package pokerHands

import (
	"fmt"
	"strings"
)

type Card struct {
	Rank Rank
	Suit Suit
}

func (c *Card) String() string {
	cardRune, ok := CardRunes[strings.ToUpper(c.Rank.Initial+c.Suit.Initial)]
	if ok {
		return fmt.Sprintf("%s %s %c", c.Rank.Initial, c.Suit.Symbol, cardRune)
	}
	return c.Rank.Initial + c.Suit.Symbol
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

var CardRunes = map[string]rune{
	"2C":  '\U0001F0D2', // 2 of Clubs
	"3C":  '\U0001F0D3', // 3 of Clubs
	"4C":  '\U0001F0D4', // 4 of Clubs
	"5C":  '\U0001F0D5', // 5 of Clubs
	"6C":  '\U0001F0D6', // 6 of Clubs
	"7C":  '\U0001F0D7', // 7 of Clubs
	"8C":  '\U0001F0D8', // 8 of Clubs
	"9C":  '\U0001F0D9', // 9 of Clubs
	"10C": '\U0001F0DA', // 10 of Clubs
	"JC":  '\U0001F0DB', // Jack of Clubs
	"QC":  '\U0001F0DD', // Queen of Clubs
	"KC":  '\U0001F0DE', // King of Clubs
	"AC":  '\U0001F0D1', // Ace of Clubs
	"2D":  '\U0001F0C2', // 2 of Diamonds
	"3D":  '\U0001F0C3', // 3 of Diamonds
	"4D":  '\U0001F0C4', // 4 of Diamonds
	"5D":  '\U0001F0C5', // 5 of Diamonds
	"6D":  '\U0001F0C6', // 6 of Diamonds
	"7D":  '\U0001F0C7', // 7 of Diamonds
	"8D":  '\U0001F0C8', // 8 of Diamonds
	"9D":  '\U0001F0C9', // 9 of Diamonds
	"10D": '\U0001F0CA', // 10 of Diamonds
	"JD":  '\U0001F0CB', // Jack of Diamonds
	"QD":  '\U0001F0CD', // Queen of Diamonds
	"KD":  '\U0001F0CE', // King of Diamonds
	"AD":  '\U0001F0C1', // Ace of Diamonds
	"2H":  '\U0001F0B2', // 2 of Hearts
	"3H":  '\U0001F0B3', // 3 of Hearts
	"4H":  '\U0001F0B4', // 4 of Hearts
	"5H":  '\U0001F0B5', // 5 of Hearts
	"6H":  '\U0001F0B6', // 6 of Hearts
	"7H":  '\U0001F0B7', // 7 of Hearts
	"8H":  '\U0001F0B8', // 8 of Hearts
	"9H":  '\U0001F0B9', // 9 of Hearts
	"10H": '\U0001F0BA', // 10 of Hearts
	"JH":  '\U0001F0BB', // Jack of Hearts
	"QH":  '\U0001F0BD', // Queen of Hearts
	"KH":  '\U0001F0BE', // King of Hearts
	"AH":  '\U0001F0B1', // Ace of Hearts
	"2S":  '\U0001F0A2', // 2 of Spades
	"3S":  '\U0001F0A3', // 3 of Spades
	"4S":  '\U0001F0A4', // 4 of Spades
	"5S":  '\U0001F0A5', // 5 of Spades
	"6S":  '\U0001F0A6', // 6 of Spades
	"7S":  '\U0001F0A7', // 7 of Spades
	"8S":  '\U0001F0A8', // 8 of Spades
	"9S":  '\U0001F0A9', // 9 of Spades
	"10S": '\U0001F0AA', // 10 of Spades
	"JS":  '\U0001F0AB', // Jack of Spades
	"QS":  '\U0001F0AD', // Queen of Spades
	"KS":  '\U0001F0AE', // King of Spades
	"AS":  '\U0001F0A1', // Ace of Spades
	"BJ":  '\U0001F0CF', // Black Joker
	"WJ":  '\U0001F0BF', // White Joker
}
