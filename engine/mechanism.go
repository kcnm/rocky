package engine

type Effect interface {
	CanHappen(game Game, you Player, target Char) bool
	Happen(game Game, you Player, target Char, cause Event)
}

type Buff interface {
	Apply(game Game, you Player, char Char)
}

type Power interface {
	Class() Class
	Mana() int
	Effect() Effect
}
