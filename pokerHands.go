package pokerHands

import (
	"slices"
	"sort"
	"strings"
)

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

type Hand struct {
	Cards []Card
	Type  string
	Value HandRank
}

func (h *Hand) String() string {
	var builder strings.Builder
	if h.Type == "" {
		h.Evaluate()
	}
	builder.WriteString(h.Type + "\n")
	for i, card := range h.Cards {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(card.String())
	}

	return builder.String()

}

func (h *Hand) Evaluate() {
	count := getCardRankCount(h.Cards)
	counterLength := len(count)
	vals := getCardValues(h.Cards)
	switch counterLength {
	case 5:
		flush := isFlush(h.Cards)
		straight := isStraight(vals)
		switch {
		case isRoyalFlush(vals, flush, straight):
			h.Type = RoyalFlushName
			h.Value = RoyalFlushRank
		case flush && straight:
			h.Type = StraightFlushName
			h.Value = StraightFlushRank
		case flush:
			h.Type = FlushName
			h.Value = FlushRank
		case straight:
			h.Type = StraightName
			h.Value = StraightRank
		default:
			h.Type = HighCardName
			h.Value = HighCardRank
		}
	case 4:
		h.Type = OnePairName
		h.Value = OnePairRank
	case 3:
		switch {
		case isThreeOfAKind(count):
			h.Type = ThreeOfAKindName
			h.Value = ThreeOfAKindRank
		case isTwoPair(count):
			h.Type = TwoPairsName
			h.Value = TwoPairsRank
		}
	case 2:
		switch {
		case isFourOfAKind(count):
			h.Type = FourOfAKindName
			h.Value = FourOfAKindRank
		case isFullHouse(count):
			h.Type = FullHouseName
			h.Value = FullHouseRank
		}
	}
}

func isOnePair(cards []Card) bool {
	counter := getCardRankCount(cards)
	if len(counter) == 4 {
		return true
	}
	return false
}

func isTwoPair(counter map[Rank]int) bool {
	if len(counter) == 3 {
		for _, c := range counter {
			if c == 2 {
				return true
			}
		}
	}
	return false
}

func isThreeOfAKind(counter map[Rank]int) bool {
	hasThreeOfAKind := false
	if len(counter) == 3 {
		for _, n := range counter {
			if n == 3 {
				hasThreeOfAKind = true
			}
		}
	}
	return hasThreeOfAKind
}

func isFourOfAKind(counter map[Rank]int) bool {
	for _, n := range counter {
		if n == 4 {
			return true
		}
	}

	return false
}

func isFullHouse(counter map[Rank]int) bool {
	hasAPair := false
	hasThree := false
	for _, n := range counter {
		if n == 2 {
			hasAPair = true
		}
		if n == 3 {
			hasThree = true
		}
	}

	return hasAPair && hasThree
}

func isRoyalFlush(cardVals []int, flush, straight bool) bool {
	if !slices.Contains(cardVals, HighAceValue) {
		return false
	}
	return flush && straight
}

func isStraight(cardVals []int) bool {
	if isConsecutive(cardVals) {
		return true
	}
	if slices.Contains(cardVals, HighAceValue) {
		for slices.Contains(cardVals, HighAceValue) {
			i := slices.Index(cardVals, HighAceValue)
			cardVals = slices.Replace(cardVals, i, i+1, LowAceValue)
		}
		return isConsecutive(cardVals)
	}
	return false
}

func isFlush(cards []Card) bool {
	suit := cards[0].Suit
	for _, card := range cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func isConsecutive(numbers []int) bool {
	// Sort the slice
	sort.Ints(numbers)
	a := numbers[0]
	b := numbers[len(numbers)-1]
	//find sum of consecutive natural numbers from a to b
	consecutiveSum := (b - a + 1) * (a + b) / 2
	//if this matches the actual sum, they are consecutive
	if consecutiveSum == sum(numbers) {
		return true
	}
	return false
}

func sum(numbers []int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}

func getCardValues(cards []Card) []int {
	cardVals := make([]int, len(cards))
	for i, card := range cards {
		cardVals[i] = card.Rank.Value
	}
	return cardVals
}

func getCardRankCount(cards []Card) map[Rank]int {
	rankCount := make(map[Rank]int)
	for _, card := range cards {
		if _, ok := rankCount[card.Rank]; !ok {
			rankCount[card.Rank] = 1
		} else {
			rankCount[card.Rank]++
		}
	}
	return rankCount
}
