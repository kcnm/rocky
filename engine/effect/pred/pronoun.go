package pred

import (
	"github.com/kcnm/rocky/engine"
)

type it struct {
	it interface{}
}

func It() Pred {
	return &it{nil}
}

func (p *it) Eval(
	game engine.Game,
	you engine.Player,
	sub interface{}) bool {
	return p.it == sub
}

func (p *it) BindIt(x interface{}) Pred {
	return &it{x}
}
