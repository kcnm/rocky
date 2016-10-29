package param

import (
	"github.com/kcnm/rocky/engine"
)

type Param interface {
	Eval(game engine.Game, you engine.Player, target engine.Char) interface{}
}
