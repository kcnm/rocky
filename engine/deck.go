package engine

import (
	"math/rand"
)

type Deck interface {
	Remain() int

	PutOnTop(card Card)
	Shuffle(rng *rand.Rand)
	Draw() (card Card, fatigue int)
}
