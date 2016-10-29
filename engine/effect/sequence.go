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

func (e sequence) CanHappen(
	game engine.Game,
	you engine.Player,
	target engine.Char) bool {
	for _, x := range e.effects {
		if !x.CanHappen(game, you, target) {
			return false
		}
	}
	return true
}

func (e sequence) Happen(
	game engine.Game,
	you engine.Player,
	target engine.Char,
	cause engine.Event) {
	events := make([]engine.Event, len(e.effects))
	for i, effect := range e.effects {
		events[i] = event.Impact(game, you, target, effect)
	}
	game.Events().Post(
		event.Sequence(game, events...), cause)
}
