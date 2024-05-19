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

	switch counterLength {
	case 5:
		switch {
		case isRoyalFlush(h.Cards):
			h.Type = RoyalFlushName
			h.Value = RoyalFlushRank
		case isStraightFlush(h.Cards):
			h.Type = StraightFlushName
			h.Value = StraightFlushRank
		case isFlush(h.Cards):
			h.Type = FlushName
			h.Value = FlushRank
		case isStraight(h.Cards):
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

func isRoyalFlush(cards []Card) bool {
	vals := getCardValues(cards)
	if !slices.Contains(vals, HighAceValue) {
		return false
	}
	return isStraightFlush(cards)
}

func isStraightFlush(cards []Card) bool {
	return isStraight(cards) && isFlush(cards)
}

func isStraight(cards []Card) bool {
	vals := getCardValues(cards)
	if isConsecutive(vals) {
		return true
	}
	if slices.Contains(vals, HighAceValue) {
		for slices.Contains(vals, HighAceValue) {
			i := slices.Index(vals, HighAceValue)
			vals = slices.Replace(vals, i, i+1, LowAceValue)
		}
		return isConsecutive(vals)
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

	// Check for consecutive values
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[i-1]+1 {
			return false
		}
	}
	return true
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
