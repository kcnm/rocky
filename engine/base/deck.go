package base

type Deck interface {
	Remain() int

	Shuffle()
	Draw() (card Card, fatigue int)
}
