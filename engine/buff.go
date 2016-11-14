package engine

type Buff interface {
	Apply(game Game, you Player, char Char)
}
