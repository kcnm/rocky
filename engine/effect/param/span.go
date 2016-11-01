package param

import (
	"github.com/kcnm/rocky/engine"
)

type span struct {
	lo, hi int
}

func Span(lo, hi int) Param {
	return &span{lo, hi}
}

func (p span) Eval(
	game engine.Game,
	you engine.Player,
	target engine.Char) interface{} {
	return p.lo + game.RNG().Intn(p.hi-p.lo+1)
}
