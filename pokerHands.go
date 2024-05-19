package pokerHands

import (
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
	return hasNOfAKind(cards, 2) == 1
}

func isTwoPair(cards []Card) bool {
	return hasNOfAKind(cards, 2) == 2
}

func isThreeOfAKind(cards []Card) bool {
	return hasNOfAKind(cards, 3) > 0
}

func isFourOfAKind(cards []Card) bool {
	return hasNOfAKind(cards, 4) > 0
}

func isFullHouse(cards []Card) bool {
	counter := getCardRankCount(cards)
	if len(counter) != 2 {
		return false
	}
	return (counter[0] == 3 && counter[1] == 2) || (counter[0] == 2 && counter[1] == 3)
}

func isRoyalFlush(cards []Card) bool {
	if !isFlush(cards) {
		return false
	}
	values := getCardValues(cards)
	for _, v := range []int{10, 11, 12, 13, 14} {
		if !contains(values, v) {
			return false
		}
	}
	return true
}

func isStraightFlush(cards []Card) bool {
	return isFlush(cards) && isStraight(cards)
}

func isStraight(cards []Card) bool {
	values := getCardValues(cards)
	sortInts(values)
	if isConsecutive(values) {
		return true
	}
	if contains(values, LowAceValue) {
		converted := convertAces(values)
		sortInts(converted)
		return isConsecutive(converted)
	}
	return false
}

func isFlush(cards []Card) bool {
	firstSuit := cards[0].Suit
	for _, card := range cards {
		if card.Suit != firstSuit {
			return false
		}
	}
	return true
}

func isConsecutive(values []int) bool {
	for i := 1; i < len(values); i++ {
		if values[i] != values[i-1]+1 {
			return false
		}
	}
	return true
}

func getCardValues(cards []Card) []int {
	values := make([]int, len(cards))
	for i, card := range cards {
		values[i] = card.Rank.Value
	}
	return values
}

func getCardRankCount(cards []Card) []int {
	counter := make(map[int]int)
	for _, card := range cards {
		counter[card.Rank.Value]++
	}
	counts := make([]int, 0, len(counter))
	for _, count := range counter {
		counts = append(counts, count)
	}
	sortInts(counts)
	return counts
}

func hasNOfAKind(cards []Card, n int) int {
	counter := make(map[int]int)
	for _, card := range cards {
		counter[card.Rank.Value]++
	}
	count := 0
	for _, v := range counter {
		if v == n {
			count++
		}
	}
	return count
}

func convertAces(values []int) []int {
	converted := make([]int, len(values))
	copy(converted, values)
	for i, v := range converted {
		if v == LowAceValue {
			converted[i] = HighAceValue
		}
	}
	return converted
}

func contains(values []int, target int) bool {
	for _, v := range values {
		if v == target {
			return true
		}
	}
	return false
}

func sortInts(values []int) {
	sort.Sort(sort.IntSlice(values))
}
