package pred

import (
	"github.com/kcnm/rocky/engine"
)

func Char(game engine.Game, you engine.Player, sub interface{}) bool {
	_, ok := sub.(engine.Char)
	return ok
}

func Hero(game engine.Game, you engine.Player, sub interface{}) bool {
	_, ok := sub.(engine.Player)
	return ok
}

func Minion(game engine.Game, you engine.Player, sub interface{}) bool {
	_, ok := sub.(engine.Minion)
	return ok
}

func Friendly(game engine.Game, you engine.Player, sub interface{}) bool {
	if !Char(game, you, sub) {
		return false
	}
	return you.IsControlling(sub.(engine.Char))
}

func Enemy(game engine.Game, you engine.Player, sub interface{}) bool {
	if !Char(game, you, sub) {
		return false
	}
	return game.Opponent(you).IsControlling(sub.(engine.Char))
}
