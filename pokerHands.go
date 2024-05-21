package pokerHands

import (
	"fmt"
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
	Cards  []Card
	Type   string
	Value  HandRank
	Used   []Card
	Unused []Card
}

func (h *Hand) String() string {
	var builder strings.Builder

	builder.WriteString(h.Type + "\n")
	for i := range h.Cards {
		if i > 0 {
			builder.WriteString(", ")
		}
		cs := fmt.Sprintf("%v", &h.Cards[i])
		builder.WriteString(cs)
	}

	return builder.String()
}

func isOnePair(ranks map[Rank]int) (bool, Rank) {
	if len(ranks) == 4 {
		for k, v := range ranks {
			if v == 2 {
				return true, k
			}
		}
	}
	return false, Rank{}
}

func isTwoPair(ranks map[Rank]int) (bool, Rank, Rank) {
	var r []Rank
	if len(ranks) == 3 {
		for k, v := range ranks {
			if v == 2 {
				r = append(r, k)
			}
		}
		if len(r) == 2 {
			return true, r[0], r[1]
		}
	}
	return false, Rank{}, Rank{}
}

func isThreeOfAKind(ranks map[Rank]int) (bool, Rank) {
	if len(ranks) == 3 {
		for k, v := range ranks {
			if v == 3 {
				return true, k
			}
		}
	}
	return false, Rank{}
}

func isFourOfAKind(ranks map[Rank]int) (bool, Rank) {
	for k, v := range ranks {
		if v == 4 {
			return true, k
		}
	}

	return false, Rank{}
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

func getRankings(cards []Card) (map[Rank][]Card, map[Rank]int, []int) {
	var cardsByRanks = make(map[Rank][]Card)
	var ranks = make(map[Rank]int)
	var vals []int
	for _, card := range cards {
		vals = append(vals, card.Rank.Value)
		if _, ok := ranks[card.Rank]; ok {
			ranks[card.Rank] = ranks[card.Rank] + 1
			cardsByRanks[card.Rank] = append(cardsByRanks[card.Rank], card)
		} else {
			ranks[card.Rank] = 1
			cardsByRanks[card.Rank] = []Card{card}
		}
	}
	return cardsByRanks, ranks, vals
}

func (h *Hand) Evaluate() {
	cardsByRanks, ranks, vals := getRankings(h.Cards)
	switch len(ranks) {
	case 5:
		flush := isFlush(h.Cards)
		straight := isStraight(vals)
		h.Used = h.Cards
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
			var highCard Card
			for i, card := range h.Cards {
				if i == 0 {
					highCard = card
				} else {
					if card.Rank.Value > highCard.Rank.Value {
						highCard = card
					}
				}
			}
			h.Used = getCardsByRank(h.Cards, highCard.Rank)
			h.Unused = getCardsByOtherRanks(cardsByRanks, highCard.Rank)
		}
	case 4:
		pair, rank := isOnePair(ranks)
		if pair {
			h.Type = OnePairName
			h.Value = OnePairRank
			h.Used = getCardsByRank(h.Cards, rank)
			h.Unused = getCardsByOtherRanks(cardsByRanks, rank)
		}

	case 3:
		three, rank := isThreeOfAKind(ranks)
		if three {
			h.Type = ThreeOfAKindName
			h.Value = ThreeOfAKindRank
			h.Used = getCardsByRank(h.Cards, rank)
			h.Unused = getCardsByOtherRanks(cardsByRanks, rank)
		} else {
			two, rnk1, rnk2 := isTwoPair(ranks)
			if two {
				h.Type = TwoPairsName
				h.Value = TwoPairsRank
				a := getCardsByRank(h.Cards, rnk1)
				b := getCardsByRank(h.Cards, rnk2)
				h.Used = append(a, b...)
				h.Unused = getCardsByOtherRanks(cardsByRanks, rnk1, rnk2)
			}
		}
	case 2:
		four, rank := isFourOfAKind(ranks)
		if four {
			h.Type = FourOfAKindName
			h.Value = FourOfAKindRank
			h.Used = getCardsByRank(h.Cards, rank)
			h.Unused = getCardsByOtherRanks(cardsByRanks, rank)
		} else {
			if isFullHouse(ranks) {
				h.Type = FullHouseName
				h.Value = FullHouseRank
				h.Used = h.Cards
			}
		}
	}
}

func getCardsByOtherRanks(cardsByRanks map[Rank][]Card, ranks ...Rank) []Card {
	returnCards := make([]Card, 0)
	for rank, cards := range cardsByRanks {
		if !slices.Contains(ranks, rank) {
			returnCards = append(returnCards, cards...)
		}
	}
	return returnCards
}

func getCardsByRank(cards []Card, rank Rank) []Card {
	returnCards := make([]Card, 0)
	for _, card := range cards {
		if card.Rank == rank {
			returnCards = append(returnCards, card)
		}
	}
	return returnCards
}

func FindWinners(hands []Hand) []Hand {
	var winners []Hand
	highestRank := 0
	for i, h := range hands {
		h.Evaluate()
		if i == 0 {
			highestRank = int(h.Value)
			winners = append(winners, h)
		} else {
			if int(h.Value) > highestRank {
				highestRank = int(h.Value)
				winners = []Hand{h}
			} else if int(h.Value) == highestRank {
				winners = append(winners, h)
			}
		}
	}
	if len(winners) == 1 {
		return winners
	}
	var realWinners []Hand
	var highUsed []int
	var highKick []int
	for i, winner := range winners {
		if i == 0 {
			realWinners = []Hand{winner}
			highUsed = getCardValues(winner.Used)
			highKick = getCardValues(winner.Unused)
			sort.Ints(highKick)
			sort.Ints(highUsed)
		} else {
			currentUsed := getCardValues(winner.Used)
			sort.Ints(currentUsed)

			switch slices.Compare(currentUsed, highUsed) {
			case -1: //highUsed is bigger
				continue
			case 0: //equal
				currentKickers := getCardValues(winner.Unused)
				sort.Ints(currentKickers)
				switch slices.Compare(currentKickers, highKick) {
				case -1: //highKick is bigger
					continue
				case 0: //equal
					realWinners = append(realWinners, winner)
					highKick = currentKickers
				case 1: //currentKickers is bigger
					highKick = currentKickers
					realWinners = []Hand{winner}
				}
			case 1: //currentUsed is bigger
				highUsed = currentUsed
				realWinners = []Hand{winner}
			}
		}
	}
	return realWinners
}
