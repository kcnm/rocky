package game

import (
	"math/rand"

	"github.com/kcnm/rocky/engine"
)

type deck struct {
	cards      []engine.Card
	fatigueCtr int
}

func NewDeck(cards ...engine.Card) engine.Deck {
	if cards == nil {
		cards = make([]engine.Card, 0)
	}
	return &deck{cards, 0}
}

func (d *deck) Remain() int {
	return len(d.cards)
}

func (d *deck) Shuffle(rng *rand.Rand) {
	shuffled := make([]engine.Card, len(d.cards))
	for i, p := range rng.Perm(len(d.cards)) {
		shuffled[p] = d.cards[i]
	}
	d.cards = shuffled
}

func (d *deck) Draw() (card engine.Card, fatigue int) {
	if len(d.cards) == 0 {
		d.fatigueCtr++
		return nil, d.fatigueCtr
	}
	card, d.cards = d.cards[0], d.cards[1:]
	return card, 0
}
