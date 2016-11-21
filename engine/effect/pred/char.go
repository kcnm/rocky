package pred

import (
	"github.com/kcnm/rocky/engine"
)

var (
	Char     = &fn{char}
	Hero     = &fn{hero}
	Minion   = &fn{minion}
	Friendly = &fn{friendly}
	Enemy    = &fn{enemy}
)

func char(game engine.Game, you engine.Player, sub interface{}) bool {
	_, ok := sub.(engine.Char)
	return ok
}

func hero(game engine.Game, you engine.Player, sub interface{}) bool {
	_, ok := sub.(engine.Player)
	return ok
}

func minion(game engine.Game, you engine.Player, sub interface{}) bool {
	_, ok := sub.(engine.Minion)
	return ok
}

func friendly(game engine.Game, you engine.Player, sub interface{}) bool {
	if !char(game, you, sub) {
		return false
	}
	return you.IsControlling(sub.(engine.Char))
}

func enemy(game engine.Game, you engine.Player, sub interface{}) bool {
	if !char(game, you, sub) {
		return false
	}
	return game.Opponent(you).IsControlling(sub.(engine.Char))
}
