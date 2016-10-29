package effect

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

type sequence struct {
	effects []engine.Effect
}

func Sequence(effects ...engine.Effect) engine.Effect {
	return &sequence{effects}
}

func (e *sequence) Happen(
	game engine.Game,
	cause engine.Event,
	target engine.Char) {
	events := make([]engine.Event, len(e.effects))
	for i, effect := range e.effects {
		events[i] = event.Impact(game, target, effect)
	}
	game.Events().Post(
		event.Sequence(game, events...), cause)
}
