package pokerHands

type HandRank int

const (
	HighCardRank HandRank = iota + 1
	OnePairRank
	TwoPairsRank
	ThreeOfAKindRank
	StraightRank
	FlushRank
	FullHouseRank
	FourOfAKindRank
	StraightFlushRank
	RoyalFlushRank
)

const (
	HighCardName      = "High Card"
	OnePairName       = "One Pair"
	TwoPairsName      = "Two Pair"
	ThreeOfAKindName  = "Three of a Kind"
	StraightName      = "Straight"
	FlushName         = "Flush"
	FullHouseName     = "Full House"
	FourOfAKindName   = "Four of a Kind"
	StraightFlushName = "Straight Flush"
	RoyalFlushName    = "Royal Flush"
)

type HandType struct {
	Rank HandRank
	Name string
}

var (
	HighCard      HandType = HandType{Rank: HighCardRank, Name: HighCardName}
	OnePair       HandType = HandType{Rank: OnePairRank, Name: OnePairName}
	TwoPairs      HandType = HandType{Rank: TwoPairsRank, Name: TwoPairsName}
	ThreeOfAKind  HandType = HandType{Rank: ThreeOfAKindRank, Name: ThreeOfAKindName}
	Straight      HandType = HandType{Rank: StraightRank, Name: StraightName}
	Flush         HandType = HandType{Rank: FlushRank, Name: FlushName}
	FullHouse     HandType = HandType{Rank: FullHouseRank, Name: FullHouseName}
	FourOfAKind   HandType = HandType{Rank: FourOfAKindRank, Name: FourOfAKindName}
	StraightFlush HandType = HandType{Rank: StraightFlushRank, Name: StraightFlushName}
	RoyalFlush    HandType = HandType{Rank: RoyalFlushRank, Name: RoyalFlushName}
)
