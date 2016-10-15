package base

type Deck interface {
	Ramain() int

	Shuffle()
	Draw() (card Card, fatigue int)
}
