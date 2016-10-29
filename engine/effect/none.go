package effect

import (
	"github.com/kcnm/rocky/engine"
)

type none int

var None none = 0

func (e none) Happen(
	game engine.Game,
	cause engine.Event,
	target engine.Char) {
	// Do nothing.
}
