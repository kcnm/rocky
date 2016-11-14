package pred

import (
	"github.com/kcnm/rocky/engine"
)

type Pred interface {
	Eval(
		game engine.Game,
		you engine.Player,
		sub interface{}) bool

	BindIt(it interface{})
}

type fn struct {
	eval func(
		game engine.Game,
		you engine.Player,
		sub interface{}) bool
}

func (p *fn) Eval(
	game engine.Game,
	you engine.Player,
	sub interface{}) bool {
	return p.eval(game, you, sub)
}

func (p *fn) BindIt(it interface{}) {}
