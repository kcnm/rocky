package impl

import (
	"math/rand"

	"github.com/kcnm/rocky/engine/base"
)

type deck struct {
	cards      []base.Card
	fatigueCtr int
}

func NewDeck(cards ...base.Card) base.Deck {
	if cards == nil {
		cards = make([]base.Card, 0)
	}
	return &deck{cards, 0}
}

func (d *deck) Remain() int {
	return len(d.cards)
}

func (d *deck) Shuffle(rng *rand.Rand) {
	shuffled := make([]base.Card, len(d.cards))
	for i, p := range rng.Perm(len(d.cards)) {
		shuffled[p] = d.cards[i]
	}
	d.cards = shuffled
}

func (d *deck) Draw() (card base.Card, fatigue int) {
	if len(d.cards) == 0 {
		d.fatigueCtr++
		return nil, d.fatigueCtr
	}
	card, d.cards = d.cards[0], d.cards[1:]
	return card, 0
}
