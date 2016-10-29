package param

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect/choose"
	"github.com/kcnm/rocky/engine/effect/pred"
)

type char struct {
	chosen choose.Choose
	pred   pred.Pred
}

func Char(chosen choose.Choose, pred pred.Pred) Param {
	return &char{chosen, pred}
}

func (p char) Eval(
	game engine.Game,
	you engine.Player,
	target engine.Char) interface{} {
	switch p.chosen {
	case choose.Manual:
		if p.pred(game, you, target) {
			return target
		} else {
			return nil
		}
	default:
		// TODO: other choose type
		return nil
	}
}
