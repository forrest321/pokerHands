package pokerHands

import (
	"math/rand"
	"slices"
)

type Deck struct {
	Cards    []Card
	Shuffled bool
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.Shuffled = false
	deck.Cards = make([]Card, 52)
	for i, card := range AllCards {
		deck.Cards[i] = card
	}
	return deck
}

func (d *Deck) RemainingCards() int {
	return len(d.Cards)
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
	if d.RemainingCards() == 0 {
		return Card{}
	}
	card := d.Cards[0]
	d.Cards = slices.Delete(d.Cards, 0, 1)
	return card
}

func (d *Deck) Deal(burn, turn int) []Card {
	cards := make([]Card, turn)
	if turn > d.RemainingCards() {
		return cards
	}
	for i := 0; i < burn; i++ {
		_ = d.DrawCard()
	}
	for i := 0; i < turn; i++ {
		cards[i] = d.DrawCard()
	}
	return cards
}
