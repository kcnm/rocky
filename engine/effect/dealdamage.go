package effect

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

type dealDamage struct {
	dmg int
}

func DealDamage(dmg int) engine.Effect {
	return &dealDamage{dmg}
}

func (e *dealDamage) Happen(
	game engine.Game,
	cause engine.Event,
	target engine.Char) {
	if target == nil {
		panic("nil target")
	}
	game.Events().Post(
		event.Damage(game, target, nil, e.dmg), cause)
}
