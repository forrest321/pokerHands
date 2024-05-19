package pokerHands

import "math/rand"

type Deck struct {
	Cards          []Card
	Discard        []Card
	RemainingCards int
	Shuffled       bool
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.RemainingCards = 52
	deck.Shuffled = false
	deck.Discard = make([]Card, 52)
	deck.Cards = AllCards
	return deck
}

func (d *Deck) Shuffle() {
	d.Shuffled = true
	for i := 0; i < 7; i++ {
		rand.Shuffle(len(d.Cards), func(i, j int) {
			d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
		})
	}
}

func (d *Deck) DrawCard() Card {
	card := d.Cards[d.RemainingCards-1]
	d.RemainingCards--
	return card
}

func (d *Deck) Deal(burn, turn int) []Card {
	var discard []Card
	var cards []Card
	for i := 0; i < burn; i++ {
		d.Discard = append(discard, d.DrawCard())
	}
	for i := 0; i < turn; i++ {
		cards = append(cards, d.DrawCard())
	}
	d.RemainingCards = d.RemainingCards - (burn + turn)
	return cards
}

var AllCards = []Card{
	{Ranks[0], Spades}, {Ranks[1], Spades}, {Ranks[2], Spades}, {Ranks[3], Spades}, {Ranks[4], Spades}, {Ranks[5], Spades}, {Ranks[6], Spades}, {Ranks[7], Spades}, {Ranks[8], Spades}, {Ranks[9], Spades}, {Ranks[10], Spades}, {Ranks[11], Spades}, {Ranks[12], Spades},
	{Ranks[0], Hearts}, {Ranks[1], Hearts}, {Ranks[2], Hearts}, {Ranks[3], Hearts}, {Ranks[4], Hearts}, {Ranks[5], Hearts}, {Ranks[6], Hearts}, {Ranks[7], Hearts}, {Ranks[8], Hearts}, {Ranks[9], Hearts}, {Ranks[10], Hearts}, {Ranks[11], Hearts}, {Ranks[12], Hearts},
	{Ranks[0], Diamonds}, {Ranks[1], Diamonds}, {Ranks[2], Diamonds}, {Ranks[3], Diamonds}, {Ranks[4], Diamonds}, {Ranks[5], Diamonds}, {Ranks[6], Diamonds}, {Ranks[7], Diamonds}, {Ranks[8], Diamonds}, {Ranks[9], Diamonds}, {Ranks[10], Diamonds}, {Ranks[11], Diamonds}, {Ranks[12], Diamonds},
	{Ranks[0], Clubs}, {Ranks[1], Clubs}, {Ranks[2], Clubs}, {Ranks[3], Clubs}, {Ranks[4], Clubs}, {Ranks[5], Clubs}, {Ranks[6], Clubs}, {Ranks[7], Clubs}, {Ranks[8], Clubs}, {Ranks[9], Clubs}, {Ranks[10], Clubs}, {Ranks[11], Clubs}, {Ranks[12], Clubs},
}
