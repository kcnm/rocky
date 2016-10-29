package engine

type Effect interface {
	Happen(game Game, cause Event, target Char)
}
