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
	rankCounts := getCardRankCount(cards)
	return len(rankCounts) == 2 && (rankCounts[0] == 3 || rankCounts[1] == 3)
}

func isRoyalFlush(cards []Card) bool {
	return containsHighAce(cards) && isStraightFlush(cards)
}

func isStraightFlush(cards []Card) bool {
	return isStraight(cards) && isFlush(cards)
}

func isStraight(cards []Card) bool {
	vals := getCardValues(cards)
	sort.Ints(vals)
	return isConsecutive(vals) || (containsLowAce(cards) && isConsecutive(convertAces(vals)))
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

func isConsecutive(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[i-1]+1 {
			return false
		}
	}
	return true
}

func getCardValues(cards []Card) []int {
	vals := make([]int, len(cards))
	for i, card := range cards {
		vals[i] = card.Rank.Value
	}
	return vals
}

func getCardRankCount(cards []Card) []int {
	counter := make(map[int]int)
	for _, card := range cards {
		counter[card.Rank.Value]++
	}
	rankCounts := make([]int, 0, len(counter))
	for _, count := range counter {
		rankCounts = append(rankCounts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rankCounts)))
	return rankCounts
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

func containsHighAce(cards []Card) bool {
	for _, card := range cards {
		if card.Rank.Value == HighAceValue {
			return true
		}
	}
	return false
}

func containsLowAce(cards []Card) bool {
	for _, card := range cards {
		if card.Rank.Value == LowAceValue {
			return true
		}
	}
	return false
}

func convertAces(vals []int) []int {
	converted := make([]int, len(vals))
	copy(converted, vals)
	for i, v := range converted {
		if v == LowAceValue {
			converted[i] = HighAceValue
		}
	}
	return converted
}
