package base

type Effect interface {
	Happen(game Game, cause Event, targets []Character)
}
