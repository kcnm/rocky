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
		if p.pred.Eval(game, you, target) {
			return target
		} else {
			return nil
		}
	case choose.All:
		chars := game.AllChars()
		result := make([]engine.Char, 0, len(chars))
		for _, ch := range chars {
			if p.pred.Eval(game, you, ch) {
				result = append(result, ch)
			}
		}
		return result
	case choose.Random:
		chars := game.AllChars()
		cands := make([]engine.Char, 0, len(chars))
		for _, ch := range chars {
			if p.pred.Eval(game, you, ch) {
				cands = append(cands, ch)
			}
		}
		result := cands[game.RNG().Intn(len(cands))]
		return result
	default:
		panic("invalid choose type")
	}
}
