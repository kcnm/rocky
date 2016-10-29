package effect

import (
	"github.com/kcnm/rocky/engine"
)

type none struct {
}

func None() engine.Effect {
	return &none{}
}

func (e *none) Happen(
	game engine.Game,
	cause engine.Event,
	targets []engine.Char) {
}
