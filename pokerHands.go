package pokerHands

import (
	"slices"
	"sort"
	"strings"
)

type HandRank int

const (
	HighCard HandRank = iota + 1
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
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
	switch {
	case isRoyalFlush(h.Cards):
		h.Type = "Royal Flush"
		h.Value = RoyalFlush
	case isStraightFlush(h.Cards):
		h.Type = "Straight Flush"
		h.Value = StraightFlush
	case isFourOfAKind(h.Cards):
		h.Type = "Four of a Kind"
		h.Value = FourOfAKind
	case isFullHouse(h.Cards):
		h.Type = "Full House"
		h.Value = FullHouse
	case isFlush(h.Cards):
		h.Type = "Flush"
		h.Value = Flush
	case isStraight(h.Cards):
		h.Type = "Straight"
		h.Value = Straight
	case isThreeOfAKind(h.Cards):
		h.Type = "Three of a Kind"
		h.Value = ThreeOfAKind
	case isTwoPair(h.Cards):
		h.Type = "Two Pair"
		h.Value = TwoPairs
	case isOnePair(h.Cards):
		h.Type = "One Pair"
		h.Value = OnePair
	default:
		h.Type = "High Card"
		h.Value = HighCard
	}
}

func isOnePair(cards []Card) bool {
	if ok, n := hasPair(cards); ok {
		return n == 1
	}
	return false
}

func isTwoPair(cards []Card) bool {
	if ok, n := hasPair(cards); ok {
		return n == 2
	}
	return false
}

func hasPair(cards []Card) (bool, int) {
	counter := getCardRankCount(cards)
	hPair := false
	pairCount := 0
	for _, n := range counter {
		if n == 2 {
			hPair = true
			pairCount++
		}
	}
	return hPair, pairCount
}

func isThreeOfAKind(cards []Card) bool {
	counter := getCardRankCount(cards)
	hasThreeOfAKind := false
	for _, n := range counter {
		if n == 3 {
			hasThreeOfAKind = true
		}
	}
	return hasThreeOfAKind
}

func isFourOfAKind(cards []Card) bool {
	counter := getCardRankCount(cards)
	for _, n := range counter {
		if n == 4 {
			return true
		}
	}

	return false
}

func isFullHouse(cards []Card) bool {
	counter := getCardRankCount(cards)
	if len(counter) == 2 {
		for _, n := range counter {
			if n == 2 || n == 3 {
				return true
			}
		}
	}
	return false
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
	if slices.Contains(vals, LowAceValue) {
		for slices.Contains(vals, LowAceValue) {
			i := slices.Index(vals, LowAceValue)
			vals = slices.Replace(vals, i, i+1, HighAceValue)
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
	var vals []int
	for _, card := range cards {
		vals = append(vals, card.Rank.Value)
	}
	return vals
}

func getCardRankCount(cards []Card) map[Rank]int {
	counter := make(map[Rank]int)
	for _, card := range cards {
		if _, ok := counter[card.Rank]; !ok {
			counter[card.Rank] = 1
		} else {
			counter[card.Rank]++
		}
	}
	return counter
}
