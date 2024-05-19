package pokerHands

type Suit struct {
	SuitName SuitName
	Initial  string
	Symbol   string
}

func (s *Suit) String() string {
	return string(s.Symbol)
}

func GetSuit(s SuitName) Suit {
	switch s {
	case SpadesName:
		return Spades
	case HeartsName:
		return Hearts
	case DiamondsName:
		return Diamonds
	case ClubsName:
		return Clubs
	default:
		return Suit{}
	}
}

type SuitName string

const (
	SpadesName   SuitName = "Spades"
	HeartsName   SuitName = "Hearts"
	DiamondsName SuitName = "Diamonds"
	ClubsName    SuitName = "Clubs"
)

var Suits = []Suit{Spades, Hearts, Diamonds, Clubs}

const (
	SpadeSymbol   = `♤`
	HeartSymbol   = `♡`
	DiamondSymbol = `♢`
	ClubSymbol    = `♧`
)

var (
	Spades   = Suit{SpadesName, "s", SpadeSymbol}
	Hearts   = Suit{HeartsName, "h", HeartSymbol}
	Diamonds = Suit{DiamondsName, "d", DiamondSymbol}
	Clubs    = Suit{ClubsName, "c", ClubSymbol}
)
