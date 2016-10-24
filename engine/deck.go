package engine

import (
	"math/rand"
)

type Deck interface {
	Remain() int

	Shuffle(rng *rand.Rand)
	Draw() (card Card, fatigue int)
}
