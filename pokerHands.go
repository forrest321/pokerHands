package pokerHands

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type Hand struct {
	Cards     Cards
	Type      HandType
	Used      Cards
	Unused    Cards
	evaluated bool
}

func (h *Hand) Len() int {
	return len(h.Cards)
}

func (h *Hand) Swap(i, j int) {
	h.Cards[i], h.Cards[j] = h.Cards[j], h.Cards[i]
}

func (h *Hand) Less(i, j int) bool {
	return h.Cards[i].Rank.Value < h.Cards[j].Rank.Value
}

type Hands []Hand

func (h *Hand) String() string {
	var builder strings.Builder

	builder.WriteString(h.Type.Name + "\n")
	for i := range h.Cards {
		if i > 0 {
			builder.WriteString(", ")
		}
		cs := fmt.Sprintf("%v", &h.Cards[i])
		builder.WriteString(cs)
	}

	return builder.String()
}

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hands) Less(i, j int) bool {
	if !h[i].evaluated {
		h[i].Evaluate()
	}
	if !h[j].evaluated {
		h[j].Evaluate()
	}
	if h[i].Type.Rank < h[j].Type.Rank {
		return true
	}
	if h[i].Type.Rank > h[j].Type.Rank {
		return false
	}
	iUsed, jUsed := getCardValues(h[i].Used), getCardValues(h[j].Used)
	switch slices.Compare(iUsed, jUsed) {
	case -1:
		return true
	case 0:
		iUnused, jUnused := getCardValues(h[i].Unused), getCardValues(h[j].Unused)
		switch slices.Compare(iUnused, jUnused) {
		case -1:
			return true
		case 0:
			return false
		case 1:
			return false
		}
	case 1:
		return false
	}
	return false
}

func FindWinners(h Hands) Hands {
	if h.Len() == 1 {
		return Hands{h[0]}
	}
	sort.Sort(h)
	var winners Hands
	topHand := h[0]
	if !topHand.evaluated {
		topHand.Evaluate()
	}
	thUsedVals := getCardValues(topHand.Used)
	for i, hh := range h {
		if !hh.evaluated {
			hh.Evaluate()
		}
		if hh.Type.Rank > topHand.Type.Rank {
			winners = Hands{hh}
			topHand = hh
			thUsedVals = getCardValues(hh.Used)
			continue
		}
		if hh.Type.Rank < topHand.Type.Rank {
			continue
		}
		if i == 0 {
			winners = Hands{topHand}
		} else {
			currentUsed := getCardValues(hh.Used)
			switch slices.Compare(thUsedVals, currentUsed) {
			case -1: //current hand wins
				topHand = hh
				thUsedVals = getCardValues(topHand.Used)
				winners = Hands{topHand}
				continue
			case 0: //used cards are equal
				// check unused
				thUnusedVals := getCardValues(topHand.Unused)
				cuUnusedVals := getCardValues(hh.Unused)
				switch slices.Compare(thUnusedVals, cuUnusedVals) {
				case -1: //current hand wins
					topHand = hh
					thUsedVals = getCardValues(topHand.Used)
					winners = Hands{topHand}
				case 0: //same
					winners = append(winners, hh)
				case 1: //topHand wins
					//no change
					continue
				}
				continue
			case 1: //topHand is bigger, no changes
				continue
			}
		}
	}

	return winners
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
	sort.Ints(cardVals)
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
	if h.evaluated {
		return
	}
	h.evaluated = true
	cardsByRanks, ranks, vals := getRankings(h.Cards)
	switch len(ranks) {
	case 5:
		flush := isFlush(h.Cards)
		straight := isStraight(vals)
		switch {
		case isRoyalFlush(vals, flush, straight):
			h.Used = h.Cards
			h.Type = RoyalFlush
		case flush && straight:
			h.Used = h.Cards
			h.Type = StraightFlush
		case flush:
			h.Used = h.Cards
			h.Type = Flush
		case straight:
			h.Used = h.Cards
			h.Type = Straight
		default:
			h.Type = HighCard
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
			h.Used = []Card{highCard}
			h.Unused = getCardsByOtherRanks(cardsByRanks, highCard.Rank)
			sort.Sort(h.Unused)
		}
	case 4:
		pair, rank := isOnePair(ranks)
		if pair {
			h.Type = OnePair
			h.Used = getCardsByRank(h.Cards, rank)
			h.Unused = getCardsByOtherRanks(cardsByRanks, rank)
			sort.Sort(h.Unused)
		}

	case 3:
		three, rank := isThreeOfAKind(ranks)
		if three {
			h.Type = ThreeOfAKind
			h.Used = getCardsByRank(h.Cards, rank)
			h.Unused = getCardsByOtherRanks(cardsByRanks, rank)
			sort.Sort(h.Unused)
		} else {
			two, rnk1, rnk2 := isTwoPair(ranks)
			if two {
				h.Type = TwoPairs
				a := getCardsByRank(h.Cards, rnk1)
				b := getCardsByRank(h.Cards, rnk2)
				h.Used = append(a, b...)
				sort.Sort(h.Used)
				h.Unused = getCardsByOtherRanks(cardsByRanks, rnk1, rnk2)
				sort.Sort(h.Unused)
			}
		}
	case 2:
		four, rank := isFourOfAKind(ranks)
		if four {
			h.Type = FourOfAKind
			h.Used = getCardsByRank(h.Cards, rank)
			h.Unused = getCardsByOtherRanks(cardsByRanks, rank)
		} else {
			if isFullHouse(ranks) {
				h.Type = FullHouse
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
