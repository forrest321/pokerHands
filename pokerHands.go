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

func (h *Hand) getCounter() map[Rank]int {
	return getCardRankCount(h.Cards)
}

func (h *Hand) Evaluate() {
	count := h.getCounter()
	counterLength := len(count)

	switch counterLength {
	case 5:
		switch {
		case isRoyalFlush(h.Cards):
			h.Type = "Royal Flush"
			h.Value = RoyalFlush
		case isStraightFlush(h.Cards):
			h.Type = "Straight Flush"
			h.Value = StraightFlush
		case isFlush(h.Cards):
			h.Type = "Flush"
			h.Value = Flush
		case isStraight(h.Cards):
			h.Type = "Straight"
			h.Value = Straight
		default:
			h.Type = "High Card"
			h.Value = HighCard
		}
	case 4:
		h.Type = "One Pair"
		h.Value = OnePair
	case 3:
		switch {
		case isThreeOfAKind(count):
			h.Type = "Three of a Kind"
			h.Value = ThreeOfAKind
		case isTwoPair(h.Cards):
			h.Type = "Two Pair"
			h.Value = TwoPairs
		}
	case 2:
		switch {
		case isFourOfAKind(count):
			h.Type = "Four of a Kind"
			h.Value = FourOfAKind
		case isFullHouse(count):
			h.Type = "Full House"
			h.Value = FullHouse
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

func isTwoPair(cards []Card) bool {
	counter := getCardRankCount(cards)
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
