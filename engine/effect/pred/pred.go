package pred

import (
	"github.com/kcnm/rocky/engine"
)

type Pred func(
	game engine.Game,
	you engine.Player,
	sub interface{}) bool

func True(game engine.Game, you engine.Player, sub interface{}) bool {
	return true
}

func False(game engine.Game, you engine.Player, sub interface{}) bool {
	return false
}

func And(preds ...Pred) Pred {
	return func(game engine.Game, you engine.Player, sub interface{}) bool {
		for _, p := range preds {
			if !p(game, you, sub) {
				return false
			}
		}
		return true
	}
}

func Or(preds ...Pred) Pred {
	return func(game engine.Game, you engine.Player, sub interface{}) bool {
		for _, p := range preds {
			if p(game, you, sub) {
				return true
			}
		}
		return false
	}
}
