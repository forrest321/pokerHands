package pokerHands

import "math/rand"

type Deck struct {
	Cards          [52]Card
	Discard        []Card
	RemainingCards int
	Shuffled       bool
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.RemainingCards = 52
	deck.Shuffled = false
	deck.Discard = make([]Card, 52)

	currentIndex := 0
	for _, suit := range Suits {
		for _, rank := range Ranks {
			deck.Cards[currentIndex] = Card{Rank: rank, Suit: suit}
			currentIndex++
		}
	}
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
