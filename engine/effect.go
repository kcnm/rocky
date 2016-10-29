package engine

type Effect interface {
	CanHappen(game Game, you Player, target Char) bool
	Happen(game Game, you Player, target Char, cause Event)
}
