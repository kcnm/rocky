package effect

import (
	"github.com/kcnm/rocky/engine"
)

type none int

var None none = 0

func (e none) CanHappen(
	game engine.Game,
	you engine.Player,
	target engine.Char) bool {
	return true
}

func (e none) Happen(
	game engine.Game,
	you engine.Player,
	target engine.Char,
	cause engine.Event) {
	// Do nothing.
}
