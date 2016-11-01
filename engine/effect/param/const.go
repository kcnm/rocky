package param

import (
	"github.com/kcnm/rocky/engine"
)

type constant struct {
	val interface{}
}

func Const(val interface{}) Param {
	return &constant{val}
}

func (p constant) Eval(
	game engine.Game,
	you engine.Player,
	target engine.Char) interface{} {
	return p.val
}
